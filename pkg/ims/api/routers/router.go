package routers

import (
	"github.com/omc-college/management-system/pkg/pubsub"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/omc-college/management-system/pkg/ims/api/handlers"
	"github.com/omc-college/management-system/pkg/ims/service"
)

//NewSignUpRouter inits Sign Up router
func NewSignUpRouter (service *service.SignUpService,client *pubsub.GroupClient) *mux.Router {

	signUpHandler := handlers.NewSignUpHandler(service,client)

	router := mux.NewRouter()


	router.HandleFunc("/signup", signUpHandler.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/email/available/{email}", signUpHandler.EmailAvailable).Methods(http.MethodGet)
	router.HandleFunc("/users/emailVerificationToken/verify/{verification_token}", signUpHandler.CheckEmailVerificationToken).Methods(http.MethodGet)
	router.HandleFunc("/reset", signUpHandler.ResetPassword).Methods(http.MethodPut)
	router.HandleFunc("/checkresettoken", signUpHandler.CheckResetToken).Methods(http.MethodGet)
	router.HandleFunc("/confirmreset", signUpHandler.ConfirmReset).Methods(http.MethodPut)
	router.HandleFunc("/sessions", ImsHandler.Login).Methods(http.MethodPost)
	router.HandleFunc("/sessions/refresh", ImsHandler.RefreshAccessToken).Methods(http.MethodPost)
	

	return router
}
