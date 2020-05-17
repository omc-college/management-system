package main

import (
	"github.com/omc-college/management-system/pkg/db"
	"github.com/omc-college/management-system/pkg/pubsub"
)

type Config struct {
	RepositoryConfig db.RepositoryConfig `mapstructure:"repository"`
	PubSubConfig pubsub.Config `mapstructure:"pubsub"`
}
