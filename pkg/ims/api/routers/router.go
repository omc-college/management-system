package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/omc-college/management-system/pkg/ims/api/handlers"
	"github.com/omc-college/management-system/pkg/ims/service"
)

//NewSignUpRouter inits Sign Up router
func NewImsRouter(service *service.ImsService) *mux.Router {

	ImsHandler := handlers.NewImsHandler(service)

	router := mux.NewRouter()

	router.HandleFunc("/signup", ImsHandler.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/email/available/{email}", ImsHandler.EmailAvailable).Methods(http.MethodGet)
	router.HandleFunc("/users/emailVerificationToken/verify/{verification_token}", ImsHandler.CheckEmailVerificationToken).Methods(http.MethodGet)
	router.HandleFunc("/sessions", ImsHandler.Login).Methods(http.MethodPost)
	router.HandleFunc("/sessions/refresh", ImsHandler.RefreshAccessToken).Methods(http.MethodPost)
	return router
}
