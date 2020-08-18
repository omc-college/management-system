package main

import (
	"github.com/omc-college/management-system/pkg/db"
	"github.com/omc-college/management-system/pkg/mq"
	"github.com/omc-college/management-system/pkg/pubsub"
)

type Config struct {
	DBConnection db.ConnectionConfig `mapstructure:"dbconnection"`
	MQConnection mq.ConnectionConfig `mapstructure:"mqconnection"`
	PubSubConfig pubsub.Config       `mapstructure:"pubsub"`
	PolicyAgent PolicyAgentConfig	`mapstructure:"policyAgent"`
}

type PolicyAgentConfig struct {
	PolicyPath	string	`mapstructure:"policyPath"`
}
