package main

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"

	"github.com/omc-college/management-system/pkg/config"
	"github.com/omc-college/management-system/pkg/pubsub"
	"github.com/omc-college/management-system/pkg/rbac"
	"github.com/omc-college/management-system/pkg/rbac/api"
	"github.com/omc-college/management-system/pkg/rbac/opa"
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

	mqURL := fmt.Sprintf("nats:%s:%s",
		serviceConfig.MQConnection.Host,
		serviceConfig.MQConnection.Port)

	client := pubsub.NewClient(serviceConfig.PubSubConfig)
	err = client.Connection(mqURL)
	if err != nil {
		logrus.Fatalf("cannot initialize Client: %s", err.Error())
	}

	subscriber := pubsub.NewQueueGroupSubscriber(*client)
	rolesChannel, err := subscriber.Subscribe(rbac.RolesTopicName)
	if err != nil {
		logrus.Fatalf("cannot subscribe on roles updates: %s", err.Error())
	}

	cache := rbac.NewCache()

	go cache.ListenUpdates(rolesChannel)

	rolesService := service.NewRolesService(repository, client)

	logrus.Fatal(http.ListenAndServe(":8000", api.NewCrudRouter(rolesService, cache, opa.GetDecision)))
}
