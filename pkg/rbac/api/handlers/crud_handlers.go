package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/omc-college/management-system/pkg/rbac/models"
	"github.com/omc-college/management-system/pkg/rbac/repository/postgres"
	"github.com/omc-college/management-system/pkg/rbac/service"
)

type RolesHandler struct {
	RolesService *service.RolesService
}

func NewRolesHandler(service *service.RolesService) *RolesHandler {
	return &RolesHandler{
		RolesService: service,
	}
}

// Handles existing error in handlers
func handleError(err error, w http.ResponseWriter) {
	var error models.Error
	var queryErr *postgres.QueryError
	var scanErr *postgres.ScanError

	if errors.As(err, &queryErr) {
		error = models.Error{http.StatusInternalServerError, queryErr.Message}
	} else if errors.As(err, &scanErr) {
		error = models.Error{http.StatusInternalServerError, queryErr.Message}
	} else if errors.Is(err, postgres.ErrCloseStmt) {
		error = models.Error{http.StatusInternalServerError, err.Error()}
	} else if errors.Is(err, postgres.ErrConvertId) {
		error = models.Error{http.StatusBadRequest, err.Error()}
	} else if errors.Is(err, postgres.ErrNoRows) {
		error = models.Error{http.StatusNotFound, err.Error()}
	} else {
		error = models.Error{http.StatusInternalServerError, err.Error()}
	}

	logrus.Errorf(error.Message)
	w.WriteHeader(error.Code)
	json.NewEncoder(w).Encode(error)
}

func (handler *RolesHandler) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error

	roles, err := handler.RolesService.GetAllRoles(r.Context())
	if err != nil {
		handleError(err, w)
		return
	}

	err = json.NewEncoder(w).Encode(roles)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *RolesHandler) GetRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error

	params := mux.Vars(r)

	if params["id"] == "" {
		err = fmt.Errorf("id is empty")
		handleError(err, w)
		return
	}

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		handleError(err, w)
		return
	}

	role, err := handler.RolesService.GetRole(r.Context(), id)
	if err != nil {
		handleError(err, w)
		return
	}

	err = json.NewEncoder(w).Encode(role)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *RolesHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var role models.Role
	var err error

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(err, w)
		return
	}

	err = json.Unmarshal(body, &role)
	if err != nil {
		handleError(err, w)
		return
	}

	err = r.Body.Close()
	if err != nil {
		handleError(err, w)
		return
	}

	err = handler.RolesService.CreateRole(r.Context(), role)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *RolesHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var role models.Role
	var err error

	params := mux.Vars(r)

	if params["id"] == "" {
		err = fmt.Errorf("id is empty")
		handleError(err, w)
		return
	}

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		handleError(err, w)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(err, w)
		return
	}

	err = json.Unmarshal(body, &role)
	if err != nil {
		handleError(err, w)
		return
	}

	err = r.Body.Close()
	if err != nil {
		handleError(err, w)
		return
	}

	err = handler.RolesService.UpdateRole(r.Context(), role, id)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *RolesHandler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error

	params := mux.Vars(r)

	if params["id"] == "" {
		err = fmt.Errorf("id is empty")
		handleError(err, w)
		return
	}

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		handleError(err, w)
		return
	}

	err = handler.RolesService.DeleteRole(r.Context(), id)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *RolesHandler) GetRoleTmpl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error

	roleTmpl, err := handler.RolesService.GetRoleTmpl(r.Context())
	if err != nil {
		handleError(err, w)
		return
	}

	err = json.NewEncoder(w).Encode(roleTmpl)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
