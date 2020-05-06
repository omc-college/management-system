package main

import (
	"flag"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/omc-college/management-system/pkg/rbac/api/routers"
	"github.com/omc-college/management-system/pkg/rbac/repository/postgres"
)

func main() {
	// Get DB config
	userFlag := flag.String("u", "postgres", "user")
	passwordFlag := flag.String("pw", "superuser", "password")
	hostFlag := flag.String("h", "localhost", "host")
	portFlag := flag.String("pt", "5432", "port")
	databaseFlag := flag.String("db", "roles", "database")
	sslmodeFlag := flag.String("ssl", "disable", "ssl mode")

	flag.Parse()

	dbConfig := fmt.Sprintf("user=%s password=%s host=%s port=%s database=%s sslmode=%s",
		*userFlag, *passwordFlag, *hostFlag, *portFlag, *databaseFlag, *sslmodeFlag)

	// Open DB
	repository, err := postgres.NewRolesRepository(dbConfig)
	if err != nil {
		log.Fatalf("opening DB error")
	}

	defer repository.DB.Close()

	// Start server
	log.Fatal(http.ListenAndServe(":8000", routers.NewCrudRouter(repository)))
}
