package main

import (
	"time"

	"github.com/omc-college/management-system/pkg/db"
)

type Config struct {
	DBConnection   db.ConnectionConfig `mapstructure:"dbconnection"`
	WebAPIAddress  string              `mapstructure:"webapi_address"`
	SigningKey     []byte              `mapstructure:"signingkey"`
	ExpirationTime time.Duration       `mapstructure:"expirationtime"`
}
