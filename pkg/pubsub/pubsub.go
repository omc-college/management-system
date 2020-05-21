package pubsub

import (
	stan "github.com/nats-io/stan.go"
)

//GroupClient contains stan.Conn
type GroupClient struct {
	stan stan.Conn
}

//NewQueueGroupClient connected to stan server and return pointer to a GroupClient
func NewQueueGroupClient(conf Config) (*GroupClient, error) {
	sc, err := stan.Connect(conf.ClusterID, conf.ClientID, stan.Pings(conf.PingsInterval, conf.MaxUnsuccessfulPings))

	return &GroupClient{
		stan: sc,
	}, err
}

//Publish publishes a message
func (stanConn *GroupClient) Publish(msg []byte, topicName string) error {
	err := stanConn.stan.Publish(topicName, msg)
	if err != nil {
		return err
	}

	return nil
}
