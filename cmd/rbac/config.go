package main

import (
	"github.com/omc-college/management-system/pkg/db"
	"github.com/omc-college/management-system/pkg/pubsub"
)

type Config struct {
	DBConnection db.ConnectionConfig `mapstructure:"dbconnection"`
	PubSubConfig pubsub.Config       `mapstructure:"pubsub"`
}
