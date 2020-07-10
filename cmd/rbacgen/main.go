package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"gopkg.in/yaml.v2"

	"github.com/omc-college/management-system/pkg/config"
	"github.com/omc-college/management-system/pkg/pubsub"
	"github.com/omc-college/management-system/pkg/rbac"
	"github.com/omc-college/management-system/pkg/rbac/repository/postgres"
	"github.com/omc-college/management-system/pkg/rbac/service"
	"github.com/omc-college/management-system/pkg/rbacgen"
)

func main() {
	configPath := flag.StringP("config", "c", "./cmd/rbacgen/rbacgen-service-example-config.yaml", "path to service config")

	isCreateMode := flag.Bool("create", false, "In this mode utility generates and creates new Role Template and saves into roleTmpl.yaml")
	isFillMode := flag.Bool("fill", false, "In this mode utility fills DB with features and endpoints from existing Role Template")

	flag.Parse()

	var serviceConfig Config

	err := config.Load(&serviceConfig, *configPath)
	if err != nil {
		logrus.Fatalf("cannot load config: %s", err.Error())
	}

	if *isCreateMode == *isFillMode {
		logrus.Fatalf("exactly one mode should be choosen")
	}

	if *isCreateMode {
		roleTmpl, err := rbacgen.GetRoleTmpl(serviceConfig.RBACGenConfig.SpecsPaths)
		if err != nil {
			logrus.Fatalf("cannot generate roleTmpl from specs: %s", err.Error())
		}

		outputFile, err := os.Create(serviceConfig.RBACGenConfig.TmplPath)
		if err != nil {
			logrus.Fatalf("cannot create new file: %s", err.Error())
		}

		roleTmplRaw, err := yaml.Marshal(roleTmpl)
		if err != nil {
			logrus.Fatalf("cannot marshal roleTmpl: %s", err.Error())
		}

		err = ioutil.WriteFile(serviceConfig.RBACGenConfig.TmplPath, roleTmplRaw, 0644)
		if err != nil {
			logrus.Fatalf("cannot write roleTmpl to a file: %s", err.Error())
		}

		err = outputFile.Close()
		if err != nil {
			logrus.Fatalf("cannot close file: %s", err.Error())
		}
	}
	if *isFillMode {
		roleTmplRaw, err := ioutil.ReadFile(serviceConfig.RBACGenConfig.TmplPath)
		if err != nil {
			logrus.Fatalf("cannot read roleTmpl from a file: %s", err.Error())
		}

		var roleTmpl rbac.RoleTmpl

		err = yaml.Unmarshal(roleTmplRaw, &roleTmpl)
		if err != nil {
			logrus.Fatalf("cannot unmarshal roleTmpl: %s", err.Error())
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

		err = repository.CreateRoleTmpl(context.Background(), roleTmpl)
		if err != nil {
			logrus.Fatalf("cannot create roleTmpl in db: %s", err.Error())
		}

		retrievedRoleTmpl, err := repository.GetRoleTmpl(context.Background())
		if err != nil {
			logrus.Fatalf("cannot get roleTmpl from db: %s", err.Error())
		}

		mqURL := fmt.Sprintf("nats://%s:%s",
			serviceConfig.MQConnection.Host,
			serviceConfig.MQConnection.Port)

		client := pubsub.NewClient(serviceConfig.PubSubConfig)
		err = client.Connection(mqURL)
		if err != nil {
			logrus.Fatalf("cannot initialize Client: %s", err.Error())
		}

		rolesService := service.NewRolesService(repository, client)

		err = rolesService.CreateRole(context.Background(), rbac.Role{
			Name:    "superuser",
			Entries: retrievedRoleTmpl.Entries,
		})
		if err != nil {
			logrus.Fatalf("cannot create superuser role in db: %s", err.Error())
		}
	}
}
