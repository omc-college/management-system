package handlers

import (
	"fmt"
	"github.com/sirupsen/logrus"

	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/omc-college/management-system/pkg/ims/validation"
	"github.com/omc-college/management-system/pkg/ims/models"
	"github.com/omc-college/management-system/pkg/ims/repository/postgresql"
	"github.com/omc-college/management-system/pkg/ims/service"
)

type SignUpHandler struct {
	SignUpService *service.SignUpService
}

func NewSignUpHandler(service *service.SignUpService) *SignUpHandler {
	return &SignUpHandler{
		SignUpService: service,
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
	} else if errors.Is(err, validation.ErrNoSymbols) || errors.Is(err, validation.ErrToMuchSymbols) {
		error = models.Error{http.StatusBadRequest, err.Error()}
	} else if errors.Is(err, validation.ErrEmailExists) || errors.Is(err, validation.ErrInvalidEmail){
		error = models.Error{http.StatusBadRequest, err.Error()}
	} else {
		error = models.Error{http.StatusInternalServerError, err.Error()}
	}

	logrus.Errorf(error.Message)
	w.WriteHeader(error.Code)
	json.NewEncoder(w).Encode(error)
}

func (h *SignUpHandler)SignUp(w http.ResponseWriter, r *http.Request) {
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

	err = validation.Data(&request)
	if err != nil {
		handleError(err, w)
		return
	}

	err = h.SignUpService.SignUp(&request)
	if err != nil {
		handleError(err, w)
		return
	}

}

func (h *SignUpHandler)EmailAvailable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error

	params := mux.Vars(r)

	if params["email"] == "" {
		err = validation.ErrNoSymbols
		handleError(err, w)
		return
	}

	result, err := h.SignUpService.EmailAvailable(params["email"])
	if err != nil {
		handleError(err, w)
		return
	} else if result == true {
		err = validation.ErrEmailExists
		handleError(err, w)
		return
	} else {
		fmt.Fprintf(w, "email is not occupied")
	}
}

func (h *SignUpHandler) CheckEmailVerificationToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tok models.EmailVerificationTokens
	var err error

	params := mux.Vars(r)

	if params["verification_token"] == "" {
		err = validation.ErrNoSymbols
		handleError(err, w)
		return
	}

	tok.VerificationToken = params["verification_token"]

	err = h.SignUpService.EmailVerificationToken(&tok)
	if err != nil {
		handleError(err, w)
		return
	}

}