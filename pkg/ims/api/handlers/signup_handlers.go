package handlers

import (
	"fmt"

	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/omc-college/management-system/pkg/ims/models"
	"github.com/omc-college/management-system/pkg/ims/repository/postgresql"
	"github.com/omc-college/management-system/pkg/hash"
)

type SignUpService struct {
	SignUpRep *postgresql.SignUpRepository
}

func (Rep *SignUpService)SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.Users
	var cred models.Credentials
	var tok models.EmailVerificationTokens

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = r.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Hash, err := hash.PwdHash(keyVal["password"])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cred.Salt = RandomGenerate(30)
	cred.PasswordHash = Hash + cred.Salt
	user.FirstName = keyVal["first_name"]
	user.LastName = keyVal["last_name"]
	user.Email = keyVal["email"]
	tok.VerificationToken = RandomGenerate(10)

	err = Rep.SignUpRep.InsertUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = Rep.SignUpRep.InsertCredentials(&cred)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = Rep.SignUpRep.InsertEmailVerificationToken(&tok)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//Send Email Verification Token function must be here!

	fmt.Fprintf(w, "User " + user.FirstName + " " + user.LastName + " registrated.")
}

func (Rep *SignUpService)EmailAvailable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emailVal string

	params := mux.Vars(r)

	result, err := Rep.SignUpRep.DB.Query("SELECT email FROM users WHERE email = $1", params["email"])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for result.Next() {
		err = result.Scan(&emailVal)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	if emailVal != "" {
		fmt.Fprintf(w, "Email already exists.")
		return
	}
}

func (Rep *SignUpService)EmailVerificationToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tok models.EmailVerificationTokens
	var user models.Users

	params := mux.Vars(r)

	tok.VerificationToken = params["verification_token"]

	result, err := Rep.SignUpRep.DB.Query("SELECT (SELECT email FROM users WHERE id = email_verification_tokens.id) FROM email_verification_tokens WHERE verification_token = $1", params["verification_token"])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for result.Next() {
		err = result.Scan(&user.Email)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	err = Rep.SignUpRep.SetUserVerified(&user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = Rep.SignUpRep.DeleteEmailVerificationToken(&tok)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Fprintf(w, "Email " + user.Email + " successfully verified. Registration completed!")
}