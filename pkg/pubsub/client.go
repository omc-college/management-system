package pubsub

import (
	stan "github.com/nats-io/stan.go"
)

//Client is a struct that contains stan.Conn and config
type Client struct {
	stan stan.Conn
	config Config
}

//NewClient is a constructor function to a Client
func NewClient(conf Config) *Client {
	return &Client{
		config: conf,
	}
}

//Connection connected to stan server
func (stanConn *Client)Connection() error {
	sc, err := stan.Connect(stanConn.config.ClusterID, stanConn.config.ClientID, stan.Pings(stanConn.config.PingsInterval, stanConn.config.MaxUnsuccessfulPings))
	if err != nil {
		return err
	}

	stanConn.stan = sc

	return nil
}

//Publish publishes a message
func (stanConn *Client) Publish(msg []byte, topicName string) error {
	err := stanConn.stan.Publish(topicName, msg)
	if err != nil {
		return err
	}

	return nil
}
