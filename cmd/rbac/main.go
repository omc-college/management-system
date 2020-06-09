package main

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"

	"github.com/omc-college/management-system/pkg/config"
	"github.com/omc-college/management-system/pkg/pubsub"
	"github.com/omc-college/management-system/pkg/rbac/api/routers"
	"github.com/omc-college/management-system/pkg/rbac/repository/postgres"
	"github.com/omc-college/management-system/pkg/rbac/service"
)

func main() {
	var serviceConfig Config
	var err error

	configPath := flag.StringP("config", "c", "./rbac-service-example-config.yaml", "path to service config")

	flag.Parse()

	err = config.Load(&serviceConfig, *configPath)
	if err != nil {
		logrus.Fatalf("cannot load config: %s", err.Error())
	}

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s database=%s sslmode=%s",
		serviceConfig.DBConnection.User, serviceConfig.DBConnection.Password, serviceConfig.DBConnection.Host,
		serviceConfig.DBConnection.Port, serviceConfig.DBConnection.Database, serviceConfig.DBConnection.Sslmode)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		logrus.Fatalf("cannot connect to db: %s", err.Error())
	}

	defer db.Close()

	repository := postgres.NewRolesRepository(db)

	client, err := pubsub.NewQueueGroupClient(serviceConfig.PubSubConfig)
	if err != nil {
		logrus.Fatalf("cannot initialize QueueGroupClient: %s", err.Error())
	}

	rolesService := service.NewRolesService(repository, client)

	logrus.Fatal(http.ListenAndServe(":8000", routers.NewCrudRouter(rolesService)))
}
