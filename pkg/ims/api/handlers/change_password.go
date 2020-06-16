package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/omc-college/management-system/pkg/ims/models"
	"github.com/omc-college/management-system/pkg/ims/service"
	"github.com/omc-college/management-system/pkg/ims/validation"
	"github.com/sirupsen/logrus"
	"net/http"
)

type CredHandler struct {
	CredService *service.CredService
}

func NewCredHandler(service *service.CredService) *CredHandler {
	return &CredHandler{
		CredService: service,
	}
}

// ErrorHandler handles existing error in handlers
func ErrorHandler(err error, w http.ResponseWriter) {
	var error models.Error

	if errors.Is(err, validation.ErrNoSymbols) || errors.Is(err, validation.ErrToMuchSymbols) {
		error = models.Error{http.StatusBadRequest, err.Error()}
	}

	logrus.Errorf(error.Message)
	w.WriteHeader(error.Code)
	json.NewEncoder(w).Encode(error)
}

func (h *CredHandler) UpdateCredentials(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var cred models.Credentials
	var err error

	params := mux.Vars(r)

	if params["newPassword "] == "" {
		err = validation.ErrNoSymbols
		ErrorHandler(err, w)
		return
	}

	cred.Password = params["newPassword"]

	err = h.CredService.ChangePassword(&cred)
	if err != nil {
		ErrorHandler(err, w)
		return
	}
}
