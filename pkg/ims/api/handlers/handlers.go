package handlers

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/omc-college/management-system/pkg/ims/models"
	"github.com/omc-college/management-system/pkg/ims/repository/postgresql"
	"github.com/omc-college/management-system/pkg/ims/service"
	"github.com/omc-college/management-system/pkg/ims/validate"
)

type ImsHandler struct {
	ImsService *service.ImsService
}

func NewImsHandler(service *service.ImsService) *ImsHandler {
	return &ImsHandler{
		ImsService: service,
	}
}

// handleError handles existing error in handlers
func handleError(err error, w http.ResponseWriter) {
	var error models.Error
	var queryErr *postgresql.QueryError
	var scanErr *postgresql.ScanError

	if errors.As(err, &queryErr) {
		error = models.Error{http.StatusInternalServerError, queryErr.Message}
	} else if errors.As(err, &scanErr) {
		error = models.Error{http.StatusInternalServerError, scanErr.Message}
	} else if errors.Is(err, validate.ErrNoSymbols) || errors.Is(err, validate.ErrToMuchSymbols) {
		error = models.Error{http.StatusBadRequest, err.Error()}
	} else if errors.Is(err, validate.ErrEmailExists) || errors.Is(err, validate.ErrInvalidEmail) {
		error = models.Error{http.StatusBadRequest, err.Error()}
	} else {
		error = models.Error{http.StatusInternalServerError, err.Error()}
	}

	logrus.Errorf(error.Message)
	w.WriteHeader(error.Code)
	json.NewEncoder(w).Encode(error)
}

func (h *ImsHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request models.SignupRequest
	var err error

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(err, w)
		return
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		handleError(err, w)
		return
	}

	err = r.Body.Close()
	if err != nil {
		handleError(err, w)
		return
	}

	err = validate.Data(&request)
	if err != nil {
		handleError(err, w)
		return
	}

	err = h.ImsService.SignUp(&request)
	if err != nil {
		handleError(err, w)
		return
	}

}

func (h *ImsHandler) EmailAvailable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error

	params := mux.Vars(r)

	if params["email"] == "" {
		err = validate.ErrNoSymbols
		handleError(err, w)
		return
	}

	result, err := h.ImsService.EmailAvailable(params["email"])
	if err != nil {
		handleError(err, w)
		return
	} else if result == true {
		err = validate.ErrEmailExists
		handleError(err, w)
		return
	} else {
		fmt.Fprintf(w, "email is not occupied")
	}
}

func (h *ImsHandler) CheckEmailVerificationToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tok models.EmailVerificationTokens
	var err error

	params := mux.Vars(r)

	if params["verification_token"] == "" {
		err = validate.ErrNoSymbols
		handleError(err, w)
		return
	}

	tok.VerificationToken = params["verification_token"]

	err = h.ImsService.EmailVerificationToken(&tok)
	if err != nil {
		handleError(err, w)
		return
	}

}

func (h *ImsHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request models.LoginRequest
	var err error

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(err, w)
		return
	}

	err = r.Body.Close()
	if err != nil {
		handleError(err, w)
		return
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		handleError(err, w)
		return
	}

	err = validate.LoginRequest(&request)
	if err != nil {
		handleError(err, w)
		return
	}

	err = h.ImsService.Login(&request)
	if err != nil {
		handleError(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ImsHandler) RefreshAccessToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error

	params := mux.Vars(r)

	if params["id"] == "" {
		err = validate.ErrNoSymbols
		handleError(err, w)
		return
	}

	err = h.ImsService.RefreshAccesssToken(params["id"])
	if err != nil {
		handleError(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ImsHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var cred models.Credentials
	var err error

	params := mux.Vars(r)

	if params["newPassword "] == "" {
		err = validate.ErrNoSymbols
		handleError(err, w)
		return
	}
	if params["newPassword "] == cred.ExistingPassword {
		err = validate.ErrConflict
		handleError(err, w)
		return
	}
	cred.ExistingPassword = params["newPassword"]

	err = h.ImsService.ChangePassword(&cred)
	if err != nil {
		handleError(err, w)
		return
	}
}
