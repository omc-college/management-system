package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/omc-college/management-system/pkg/ims/api/handlers"
	"github.com/omc-college/management-system/pkg/ims/repository/postgresql"
)

//NewSignUpRouter inits Sign Up router
func NewSignUpRouter (repository *postgresql.SignUpRepository) *mux.Router {

	var signUpHandler handlers.SignUpService
	signUpHandler.SignUpRep = repository

	router := mux.NewRouter()

	router.HandleFunc("/signup", signUpHandler.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/email/available/{email}", signUpHandler.EmailAvailable).Methods(http.MethodGet)
	router.HandleFunc("/users/emailVerificationToken/verify/{verification_token}", signUpHandler.EmailVerificationToken).Methods(http.MethodGet)

	return router
}