package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

//jwt token
var mySigningKey = []byte("secret")
var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["admin"] = true
	claims["name"] = "dafdgdsds"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningKey)

	/* Finally, write the token to the browser window */
	w.Write([]byte(tokenString))
})

func main() {
	//Init Router
	r := mux.NewRouter()

	r.Handle("/get-token", GetTokenHandler).Methods("GET")


	log.Fatal(http.ListenAndServe(":8001", r))
}
