package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"

	"github.com/gorilla/handlers"

	"github.com/omc-college/management-system/pkg/ims/api/routers"
	"github.com/omc-college/management-system/pkg/ims/service"
	"github.com/omc-college/management-system/pkg/config"
)

func main() {
	var conf Config
	var err error

	configPath := flag.StringP("config", "c", "./ims-service-example-config.yaml", "path to service config")

	flag.Parse()

	err = config.Load(&conf, *configPath)
	if err != nil {
		logrus.Fatalf("cannot load config: %s", err.Error())
	}

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s database=%s sslmode=%s",
		conf.DBConnection.User, conf.DBConnection.Password, conf.DBConnection.Host,
		conf.DBConnection.Port, conf.DBConnection.Database, conf.DBConnection.Sslmode)

	// Open DB
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		logrus.Fatalf("cannot connect to DB: %s", err.Error())
	}

	defer db.Close()

	signupService := service.NewIMSService(db, nil, time.Time{})

	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "DELETE", "POST", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})

	// Start server
	logrus.Fatal(http.ListenAndServe(conf.WebAPIAddress, handlers.CORS(headers, methods, origins)(routers.NewImsRouter(signupService))))
}
