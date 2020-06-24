package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/omc-college/management-system/pkg/config"
	"github.com/omc-college/management-system/pkg/ri/api"
	"github.com/omc-college/management-system/pkg/ri/repository/postgresql"
	"github.com/omc-college/management-system/pkg/ri/service"
	"net/http"

	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	//"net/http"
)

func main() {
	var serviceConfig Config
	var err error

	configPath := flag.StringP("config", "c", "./cmd/ri/ri-config.yaml", "path to service config")

	flag.Parse()

	err = config.Load(&serviceConfig, *configPath)
	if err != nil {
		logrus.Fatalf("%s", err)
	}

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s database=%s sslmode=%s",
		serviceConfig.RepositoryConfig.User, serviceConfig.RepositoryConfig.Password, serviceConfig.RepositoryConfig.Host, serviceConfig.RepositoryConfig.Port, serviceConfig.RepositoryConfig.Database, serviceConfig.RepositoryConfig.Sslmode)

	db, err := sqlx.Connect("pgx", dsn)

	if err != nil {
		logrus.Fatalf("opening DB error")
	}

	defer db.Close()

	repository := postgresql.NewResourcesRepository(db)

	service := service.NewResourcesService (repository)

	http.ListenAndServe(":8000", api.NewResourcesRouter(service))



}