package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/omc-college/management-system/pkg/rbac"
	"github.com/omc-college/management-system/pkg/rbac/service"
)

// NewCrudRouter Inits RBAC CRUD service router
func NewCrudRouter(service *service.RolesService, cache *rbac.Cache, decide func(context.Context, rbac.Input) error) *mux.Router {
	authorizationMiddleware := rbac.NewRBACMiddleware(cache, decide)
	rolesHandler := NewRolesHandler(service)

	router := mux.NewRouter()

	router.HandleFunc("/roles", rolesHandler.GetAllRoles).Methods(http.MethodGet)
	router.HandleFunc("/roles", rolesHandler.CreateRole).Methods(http.MethodPost)
	router.HandleFunc("/roles/{id}", rolesHandler.GetRole).Methods(http.MethodGet)
	router.HandleFunc("/roles/{id}", rolesHandler.UpdateRole).Methods(http.MethodPut)
	router.HandleFunc("/roles/{id}", rolesHandler.DeleteRole).Methods(http.MethodDelete)
	router.HandleFunc("/roletmpl", rolesHandler.GetRoleTmpl).Methods(http.MethodGet)

	router.Use(authorizationMiddleware.Middleware)

	return router
}
