package main

import (
	"github.com/omc-college/management-system/pkg/db"
)

type Config struct {
	RepositoryConfig db.RepositoryConfig `mapstructure:"repository"`
}
