package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/omc-college/management-system/pkg/ims/api/handlers"
	"github.com/omc-college/management-system/pkg/ims/service"
)

//NewSignUpRouter inits Sign Up router
func NewImsRouter(service *service.ImsService) *mux.Router {

	imsHandler := handlers.NewImsHandler(service)

	router := mux.NewRouter()

	router.HandleFunc("/sign-up", imsHandler.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/email/available/{email}", imsHandler.EmailAvailable).Methods(http.MethodGet)
	router.HandleFunc("/users/emailVerificationToken/verify/{verification_token}", imsHandler.CheckEmailVerificationToken).Methods(http.MethodGet)
	router.HandleFunc("/sessions", imsHandler.Login).Methods(http.MethodPost)
	router.HandleFunc("/sessions/refresh", imsHandler.RefreshAccessToken).Methods(http.MethodPost)
	router.HandleFunc("/password", imsHandler.ChangePassword).Methods(http.MethodPost)

	return router
}
