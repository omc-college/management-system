package main

import (
	"github.com/omc-college/management-system/pkg/db"
)

type Config struct {
	DBConnection  db.ConnectionConfig `mapstructure:"dbconnection"`
	RBACGenConfig RBACGenConfig       `mapstructure:"rbacgen"`
}

type RBACGenConfig struct {
	SpecsPaths []string `mapstructure:"specsPaths"`
	TmplPath   string   `mapstructure:"tmplPath"`
}
