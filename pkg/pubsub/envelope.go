package pubsub

import (
	"encoding/json"

	"github.com/nats-io/stan.go"
)

type Envelope interface {
	EntityType() string
	Operation() string
	Payload() json.RawMessage
	Ack() error
}

type envelope struct {
	entityType string
	operation string
	payload json.RawMessage
	msg *stan.Msg
}

func NewEnvelope(payload interface{}, operation string, entityType string) (*envelope, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return &envelope{
		entityType: entityType,
		operation: operation,
		payload:   payloadBytes,
	}, err
}

func (e *envelope)Ack() error {
	err := e.msg.Ack()
	if err != nil {
		return err
	}

	return nil
}

func (e *envelope)EntityType() string {
	return e.entityType
}

func (e *envelope)Operation() string {
	return e.operation
}

func (e *envelope)Payload() json.RawMessage {
	return e.payload
}