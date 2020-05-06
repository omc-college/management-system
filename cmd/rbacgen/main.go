package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"gopkg.in/yaml.v2"

	"github.com/omc-college/management-system/pkg/rbac/models"
	"github.com/omc-college/management-system/pkg/rbac/openapi"
	"github.com/omc-college/management-system/pkg/rbac/repository/postgres"
)

func main() {
	var roleTmpl models.RoleTmpl
	var roleTmplRaw []byte

	// Get DB config
	userFlag := flag.String("u", "postgres", "user")
	passwordFlag := flag.String("pw", "superuser", "password")
	hostFlag := flag.String("h", "localhost", "host")
	portFlag := flag.String("pt", "5432", "port")
	databaseFlag := flag.String("db", "roles", "database")
	sslmodeFlag := flag.String("ssl", "disable", "ssl mode")

	// Get specs path's and output path
	specs := flag.StringSlice("specs", nil, "Path to specs that should be parsed")
	tmpl := flag.String("tmpl", "", "Path where to save RBAC template")

	// Get mode
	isCreateMode := flag.Bool("create", false, "In this mode utility generates and creates new Role Template and saves into roleTmpl.yaml")
	isFillMode := flag.Bool("fill", false, "In this mode utility fills DB with features and endpoints from existing Role Template")

	flag.Parse()

	// check whether only one mode flag is provided
	if *isCreateMode == *isFillMode {
		err := fmt.Errorf("you should choose only one mode")
		logrus.Fatalf("%s", err)
	}

	if *isCreateMode {
		// generate and get new Role Template
		roleTmpl, err := openapi.GetRoleTmpl(*specs)
		if err != nil {
			logrus.Fatalf("%s", err)
		}

		// Creating output file
		outputFile, err := os.Create(*tmpl)
		if err != nil {
			logrus.Fatalf("%s", err)
		}

		// Encoding and writing YAML into new file
		roleTmplRaw, err = yaml.Marshal(roleTmpl)
		if err != nil {
			logrus.Fatalf("%s", err)
		}

		err = ioutil.WriteFile(*tmpl, roleTmplRaw, 0644)
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
		roleTmplRaw, err := ioutil.ReadFile(*tmpl)
		if err != nil {
			logrus.Fatalf("%s", err)
		}

		yaml.Unmarshal(roleTmplRaw, &roleTmpl)

		// Form db config
		dbConfig := fmt.Sprintf("user=%s password=%s host=%s port=%s database=%s sslmode=%s",
			*userFlag, *passwordFlag, *hostFlag, *portFlag, *databaseFlag, *sslmodeFlag)

		// Open DB
		repository, err := postgres.NewRolesRepository(dbConfig)
		if err != nil {
			logrus.Fatalf("opening DB error")
		}
		defer repository.DB.Close()

		err = postgres.CreateRoleTmpl(repository, roleTmpl)
		if err != nil {
			logrus.Fatalf("%s", err)
		}
	}
}
