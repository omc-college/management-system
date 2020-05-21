package pubsub

import (
	"encoding/json"
)

type Envelope struct {
	Type string
	Operation string
	Payload json.RawMessage
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

