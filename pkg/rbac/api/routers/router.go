package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	handlers "github.com/omc-college/management-system/pkg/rbac/api/handlers"
	postgres "github.com/omc-college/management-system/pkg/rbac/repository/postgres"
)

// NewCrudRouter Inits RBAC CRUD service router
func NewCrudRouter(repository *postgres.RolesRepository) *mux.Router {
	// Init handlers DB wrap
	var rolesHandler handlers.RolesHandler
	rolesHandler.RolesRepository = repository

	router := mux.NewRouter()

	router.HandleFunc("/roles", rolesHandler.GetAllRoles).Methods(http.MethodGet)
	router.HandleFunc("/roles", rolesHandler.CreateRole).Methods(http.MethodPost)
	router.HandleFunc("/roles/{id}", rolesHandler.GetRole).Methods(http.MethodGet)
	router.HandleFunc("/roles/{id}", rolesHandler.UpdateRole).Methods(http.MethodPut)
	router.HandleFunc("/roles/{id}", rolesHandler.DeleteRole).Methods(http.MethodDelete)
	router.HandleFunc("/roletmpl", rolesHandler.GetRoleTemplate).Methods(http.MethodGet)

	return router
}
