package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"gopkg.in/yaml.v2"

	"github.com/omc-college/management-system/pkg/config"
	"github.com/omc-college/management-system/pkg/rbac/models"
	"github.com/omc-college/management-system/pkg/rbac/openapi"
	"github.com/omc-college/management-system/pkg/rbac/repository/postgres"
)

func main() {
	var roleTmpl models.RoleTmpl
	var roleTmplRaw []byte
	var serviceConfig Config
	var err error

	ctx := context.Background()

	configPath := flag.StringP("config", "c", "cmd/rbacgen/rbacgen-service-example-config.yaml", "path to service config")

	// Get mode
	isCreateMode := flag.Bool("create", false, "In this mode utility generates and creates new Role Template and saves into roleTmpl.yaml")
	isFillMode := flag.Bool("fill", false, "In this mode utility fills DB with features and endpoints from existing Role Template")

	flag.Parse()

	err = config.Load(&serviceConfig, *configPath)
	if err != nil {
		logrus.Fatalf("%s", err)
	}

	// check whether only one mode flag is provided
	if *isCreateMode == *isFillMode {
		err := fmt.Errorf("you should choose only one mode")
		logrus.Fatalf("%s", err)
	}

	if *isCreateMode {
		// generate and get new Role Template
		roleTmpl, err := openapi.GetRoleTmpl(serviceConfig.RBACGenConfig.SpecsPaths)
		if err != nil {
			logrus.Fatalf("%s", err)
		}

		// Creating output file
		outputFile, err := os.Create(serviceConfig.RBACGenConfig.TmplPath)
		if err != nil {
			logrus.Fatalf("%s", err)
		}

		// Encoding and writing YAML into new file
		roleTmplRaw, err = yaml.Marshal(roleTmpl)
		if err != nil {
			logrus.Fatalf("%s", err)
		}

		err = ioutil.WriteFile(serviceConfig.RBACGenConfig.TmplPath, roleTmplRaw, 0644)
		if err != nil {
			logrus.Fatalf("%s", err)
		}

		// Closing new file
		err = outputFile.Close()
		if err != nil {
			logrus.Fatalf("%s", err)
		}
	}
	if *isFillMode {
		// get existing Role Template from file
		roleTmplRaw, err := ioutil.ReadFile(serviceConfig.RBACGenConfig.TmplPath)
		if err != nil {
			logrus.Fatalf("%s", err)
		}

		yaml.Unmarshal(roleTmplRaw, &roleTmpl)

		// Open DB
		repository, err := postgres.NewRolesRepository(serviceConfig.RepositoryConfig)
		if err != nil {
			logrus.Fatalf("opening DB error")
		}
		defer repository.DB.Close()

		err = repository.CreateRoleTmpl(ctx, roleTmpl)
		if err != nil {
			logrus.Fatalf("%s", err)
		}
	}
}
