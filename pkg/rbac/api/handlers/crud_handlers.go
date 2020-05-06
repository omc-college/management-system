package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/omc-college/management-system/pkg/rbac/models"
	"github.com/omc-college/management-system/pkg/rbac/repository/postgres"
)

type RolesHandler struct {
	RolesRepository *postgres.RolesRepository
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

	log.Errorf(error.Message)
	w.WriteHeader(error.Code)
	json.NewEncoder(w).Encode(error)
}

func (repository *RolesHandler) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	roles, err := postgres.GetAllRoles(repository.RolesRepository)
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

func (repository *RolesHandler) GetRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	if params["id"] == "" {
		err := fmt.Errorf("id is empty")
		handleError(err, w)
		return
	}

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		handleError(fmt.Errorf("converting id to int error"), w)
		return
	}

	role, err := postgres.GetRole(repository.RolesRepository, id)
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

func (repository *RolesHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(err, w)
		return
	}

	var role models.Role

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

	err = postgres.CreateRole(repository.RolesRepository, role)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (repository *RolesHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	if params["id"] == "" {
		err := fmt.Errorf("id is empty")
		handleError(err, w)
		return
	}

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		handleError(fmt.Errorf("converting id to int error"), w)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(err, w)
		return
	}

	var role models.Role
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

	err = postgres.UpdateRole(repository.RolesRepository, role, id)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (repository *RolesHandler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	if params["id"] == "" {
		err := fmt.Errorf("id is empty")
		handleError(err, w)
		return
	}

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		handleError(fmt.Errorf("converting id to int error"), w)
		return
	}

	err = postgres.DeleteRole(repository.RolesRepository, id)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (repository *RolesHandler) GetRoleTemplate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	roleTemplate, err := postgres.GetRoleTmpl(repository.RolesRepository)
	if err != nil {
		handleError(err, w)
		return
	}

	err = json.NewEncoder(w).Encode(roleTemplate)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
