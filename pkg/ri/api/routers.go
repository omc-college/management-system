package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/omc-college/management-system/pkg/ri/service"
)

// NewResourcesRouter Inits RI CRUD service router
func NewResourcesRouter(service *service.ResourcesService) *mux.Router {

	resourcesHandler := NewResourcesHandler(service)

	router := mux.NewRouter()

	router.HandleFunc("/resources", resourcesHandler.GetAllResources).Methods(http.MethodGet)
	router.HandleFunc("/resources", resourcesHandler.InsertResource).Methods(http.MethodPost)
	router.HandleFunc("/resources/{id}", resourcesHandler.GetResource).Methods(http.MethodGet)
	router.HandleFunc("/resources/{id}", resourcesHandler.UpdateResource).Methods(http.MethodPut)
	router.HandleFunc("/resources/{id}", resourcesHandler.DeleteResource).Methods(http.MethodDelete)

	return router
}
