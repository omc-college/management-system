package main

import (
	"github.com/omc-college/management-system/pkg/db"
	"github.com/omc-college/management-system/pkg/pubsub"
)

type Config struct {
	DBConnection  db.ConnectionConfig `mapstructure:"dbconnection"`
	RBACGenConfig RBACGenConfig       `mapstructure:"rbacgen"`
	PubSubConfig  pubsub.Config       `mapstructure:"pubsub"`
}

type RBACGenConfig struct {
	SpecsPaths []string `mapstructure:"specsPaths"`
	TmplPath   string   `mapstructure:"tmplPath"`
}
