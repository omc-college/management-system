package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/omc-college/management-system/pkg/ims/api/handlers"
	"github.com/omc-college/management-system/pkg/ims/service"
)

//NewSignUpRouter inits Sign Up router
func NewSignUpRouter(service *service.SignUpService) *mux.Router {

	signUpHandler := handlers.NewSignUpHandler(service)

	router := mux.NewRouter()

	router.HandleFunc("/signup", signUpHandler.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/email/available/{email}", signUpHandler.EmailAvailable).Methods(http.MethodGet)
	router.HandleFunc("/users/emailVerificationToken/verify/{verification_token}", signUpHandler.CheckEmailVerificationToken).Methods(http.MethodGet)

	return router
}

//NewCredRouter inits Change Password router
func NewCredRouter(service *service.CredService) *mux.Router {

	CredHandler := handlers.NewCredHandler(service)

	router := mux.NewRouter()

	router.HandleFunc("/password", CredHandler.UpdateCredentials).Methods(http.MethodPost)

	return router
}
