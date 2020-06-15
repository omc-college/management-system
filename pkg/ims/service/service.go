package service

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"

	"github.com/omc-college/management-system/pkg/ims/models"
	"github.com/omc-college/management-system/pkg/ims/repository/postgresql"
	tokenCreate "github.com/omc-college/management-system/pkg/jwt"
	"github.com/omc-college/management-system/pkg/pwd"
)

// token returns random int for verification token
func token() string {
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(999999999999999)))
	rand := strconv.Itoa(int(r.Int64()))
	return rand
}

type ImsService struct {
	db             *sqlx.DB
	signingKey     []byte
	expirationTime time.Time
}

func NewIMSService(DB *sqlx.DB, signingKey []byte, expirationTime time.Time) *ImsService {
	return &ImsService{
		db:             DB,
		signingKey:     signingKey,
		expirationTime: expirationTime,
	}
}

func (service *ImsService) SignUp(request *models.SignupRequest) error {
	var cred models.Credentials
	var tok models.EmailVerificationTokens
	var err error
	var user models.User

	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email

	tok.VerificationToken = token()

	cred.Salt = pwd.Salt(256 - len(request.Password))

	cred.PasswordHash, err = pwd.Hash(request.Password, cred.Salt)
	if err != nil {
		return err
	}

	tx, err := service.db.Beginx()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	signUpRepo := postgresql.NewSignUpRepository(tx)
	credRepo := postgresql.NewCredentialsRepository(tx)
	userRepo := postgresql.NewUsersRepository(tx)

	err = userRepo.InsertUser(&user)
	if err != nil {
		return err
	}

	tok.ID = user.ID
	cred.ID = user.ID

	err = credRepo.InsertCredentials(&cred)
	if err != nil {
		return err
	}

	err = signUpRepo.InsertEmailVerificationToken(&tok)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (service *ImsService) EmailAvailable(email string) (bool, error) {
	var user *models.User
	var exist bool
	var err error

	userRepo := postgresql.NewUsersRepository(service.db)

	user, err = userRepo.GetUserByEmail(email)
	if err != nil {
		return exist, err
	}

	if user.Email != "" {
		exist = true
	}

	return exist, nil
}

func (service *ImsService) EmailVerificationToken(token *models.EmailVerificationTokens) error {
	var user models.User
	var err error

	tx, err := service.db.Beginx()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	userRepo := postgresql.NewUsersRepository(tx)
	signUpRepo := postgresql.NewSignUpRepository(tx)

	user.Email, err = userRepo.GetEmailByToken(token)
	if err != nil {
		return err
	}

	err = signUpRepo.SetUserVerified(&user)
	if err != nil {
		return err
	}

	err = signUpRepo.DeleteEmailVerificationToken(token)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Send Email Verification Token function must be here!

func (service *ImsService) Login(request *models.LoginRequest) error {
	var user *models.User
	var cred *models.Credentials
	var err error
	var id, ss string
	userRepo := postgresql.NewUsersRepository(service.db)
	credRepo := postgresql.NewCredentialsRepository(service.db)
	user, err = userRepo.GetUserByEmail(request.Email)
	if err != nil {
		return err
	}

	id = strconv.Itoa(user.ID)
	cred, err = credRepo.GetCredentialByUserID(id)
	if err != nil {
		return err
	}
	pwd := []byte(request.Password + cred.Salt)
	hashedPasword := []byte(cred.PasswordHash)
	err = bcrypt.CompareHashAndPassword(hashedPasword, pwd)
	if err != nil {
		return err
	}

	claims := tokenCreate.Claims{
		id,
		user.FirstName,
		user.Email,
		user.Roles,
		jwt.StandardClaims{
			ExpiresAt: service.expirationTime.Unix(),
		},
	}

	ss, err = tokenCreate.GenerateToken(claims, service.signingKey)
	if err != nil {
		return err
	}
	err = credRepo.InsertAccessToken(id, ss)
	if err != nil {
		return err
	}
	return nil
}

func (service *ImsService) RefreshAccesssToken(id string) error {
	var user *models.User
	var err error
	var ss string

	userRepo := postgresql.NewUsersRepository(service.db)
	credRepo := postgresql.NewCredentialsRepository(service.db)
	user, err = userRepo.GetUserByEmail(id)
	if err != nil {
		return err
	}

	claims := tokenCreate.Claims{
		id,
		user.FirstName,
		user.Email,
		user.Roles,
		jwt.StandardClaims{
			ExpiresAt: service.expirationTime.Unix(),
		},
	}
	ss, err = tokenCreate.GenerateToken(claims, service.signingKey)
	if err != nil {
		return err
	}
	err = credRepo.UpdateAccessToken(id, ss)
	if err != nil {
		return err
	}
	return nil
}
