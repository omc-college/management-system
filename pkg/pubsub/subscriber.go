package pubsub

import (
	"encoding/json"

	stan "github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

type GroupSubscriber struct {
	stanSubscriber Client
	subscription   stan.Subscription
}

//NewQueueGroupSubscriber is a constructor function to a GroupSubscriber
func NewQueueGroupSubscriber(client Client) *GroupSubscriber {
	return &GroupSubscriber{
		stanSubscriber: client,
	}
}

//Subscribe put receive message to the channel
func (stanConn *GroupSubscriber) Subscribe(topicName string) (<-chan Envelope, error) {
	output := make(chan Envelope)

	sub, err := stanConn.subscribe(output, topicName)
	if err != nil {
		return nil, err
	}

	stanConn.subscription = sub

	return output, nil
}

func (stanConn *GroupSubscriber) subscribe(output chan Envelope, topicName string) (stan.Subscription, error) {
	return stanConn.stanSubscriber.stan.Subscribe(topicName, func(m *stan.Msg) {
		var e = &envelope{}

		err := json.Unmarshal(m.Data, e)
		if err != nil {
			logrus.Error(err.Error())
			return
		}

		e.Msg = m

		output <- e

		m.Ack()
	}, stan.SetManualAckMode(), stan.DeliverAllAvailable())
}

//Unsubscribe removes interest in the subscription
func (stanConn *GroupSubscriber) Unsubscribe() error {
	err := stanConn.subscription.Unsubscribe()
	if err != nil {
		return err
	}

	return nil
}

//Close removes subscriber from the server
func (stanConn *GroupSubscriber) Close() error {
	err := stanConn.subscription.Close()
	if err != nil {
		return err
	}

	return nil
}
