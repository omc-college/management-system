package pubsub

import (
	"encoding/json"

	"github.com/nats-io/stan.go"
)

type Envelope interface {
	GetEntityType() string
	GetOperation() string
	GetPayload() json.RawMessage
	Ack() error
}

type envelope struct {
	EntityType string
	Operation  string
	Payload    json.RawMessage
	Msg        *stan.Msg
}

func NewEnvelope(payload interface{}, operation string, entityType string) (*envelope, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return &envelope{
		EntityType: entityType,
		Operation:  operation,
		Payload:    payloadBytes,
	}, err
}

func (e *envelope) Ack() error {
	err := e.Msg.Ack()
	if err != nil {
		return err
	}

	return nil
}

func (e *envelope) GetEntityType() string {
	return e.EntityType
}

func (e *envelope) GetOperation() string {
	return e.Operation
}

func (e *envelope) GetPayload() json.RawMessage {
	return e.Payload
}
