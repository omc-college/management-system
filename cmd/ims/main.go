package main

import (
	"github.com/omc-college/management-system/pkg/ims/api/routers"
	"github.com/omc-college/management-system/pkg/ims/service"
	"log"
	"net/http"
)

func main() {
	routers.NewSignUpRouter(&service.SignUpService{})
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
