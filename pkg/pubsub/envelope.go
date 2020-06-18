package pubsub

import (
	"encoding/json"

	"github.com/nats-io/stan.go"
)

type Envelope struct {
	Type string
	Operation string
	Payload json.RawMessage
	msg *stan.Msg
}

func NewEnvelope(payload interface{}, operation string, entityType string) (*Envelope, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return &Envelope{
		Type:      entityType,
		Operation: operation,
		Payload:   payloadBytes,
	}, err
}

func (e *Envelope)Ack() error {
	err := e.msg.Ack()
	if err != nil {
		return err
	}

	return nil
}

