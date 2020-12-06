package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"errors"
	"strconv"

	"github.com/omc-college/management-system/pkg/ri/models"
	"github.com/omc-college/management-system/pkg/ri/repository/postgresql"
	"github.com/omc-college/management-system/pkg/ri/service"
	"github.com/sirupsen/logrus"

)

type ResourcesHandler struct {
	ResourcesService *service.ResourcesService
}

func NewResourcesHandler(service *service.ResourcesService) *ResourcesHandler {
	return &ResourcesHandler{
		ResourcesService: service,
	}
}

func handleError(err error, w http.ResponseWriter) {
	var error models.Error
	var queryErr *postgresql.QueryError
	var scanErr *postgresql.ScanError

	if errors.As(err, &queryErr) {
		error = models.Error{http.StatusInternalServerError, queryErr.Message}
	} else if errors.As(err, &scanErr) {
		error = models.Error{http.StatusInternalServerError, queryErr.Message}
	} else if errors.Is(err, postgresql.ErrCloseStmt) {
		error = models.Error{http.StatusInternalServerError, err.Error()}
	} else if errors.Is(err, postgresql.ErrConvertId) {
		error = models.Error{http.StatusBadRequest, err.Error()}
	} else if errors.Is(err, postgresql.ErrNoRows) {
		error = models.Error{http.StatusNotFound, err.Error()}
	} else {
		error = models.Error{http.StatusInternalServerError, err.Error()}
	}

	logrus.Errorf(error.Message)
	w.WriteHeader(error.Code)
	json.NewEncoder(w).Encode(error)
}

func (handler *ResourcesHandler) GetAllResources(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error

	resources, err := handler.ResourcesService.GetAllResources()
	if err != nil {
		handleError(err, w)
		return
	}

	err = json.NewEncoder(w).Encode(resources)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *ResourcesHandler) GetResource(w http.ResponseWriter, r *http.Request) {
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

	resource, err := handler.ResourcesService.GetResource(id)
	if err != nil {
		handleError(err, w)
		return
	}

	err = json.NewEncoder(w).Encode(resource)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *ResourcesHandler) InsertResource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resource models.Resources
	var err error

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(err, w)
		return
	}

	err = json.Unmarshal(body, &resource)
	if err != nil {
		handleError(err, w)
		return
	}

	err = r.Body.Close()
	if err != nil {
		handleError(err, w)
		return
	}

	err = handler.ResourcesService.InsertResource(resource)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *ResourcesHandler) UpdateResource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resource models.Resources
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

	err = json.Unmarshal(body, &resource)
	if err != nil {
		handleError(err, w)
		return
	}

	err = r.Body.Close()
	if err != nil {
		handleError(err, w)
		return
	}

	err = handler.ResourcesService.UpdateResource(resource, id)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *ResourcesHandler) DeleteResource(w http.ResponseWriter, r *http.Request) {
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

	err = handler.ResourcesService.DeleteResource(id)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
