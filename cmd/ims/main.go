package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/omc-college/management-system/pkg/ims/api/routers"
	"github.com/omc-college/management-system/pkg/ims/service"
)

type connection struct {
	user string;
	password string;
	host string;
	port string;
	database string;
	sslmode string;
}

func main() {
	var err error
	var conn connection

	// Mock data
	conn.user = "postgres";
	conn.password = "krimyllexadmin";
	conn.host = "localhost";
	conn.port = "5432";
	conn.database = "ims";
	conn.sslmode = "disable";

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s database=%s sslmode=%s",
		conn.user, conn.password, conn.host, conn.port, conn.database, conn.sslmode)

	// Open DB
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		logrus.Fatalf("cannot connect to DB: %s", err.Error())
	}

	defer db.Close()

	signupService := service.NewSignUpService(db)

	// Start server
	logrus.Fatal(http.ListenAndServe(":8080", routers.NewSignUpRouter(signupService)))

}
