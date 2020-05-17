package pubsub

type Config struct {
	ClusterID string
	ClientID string
	QueueName string
	PingsInterval int
	MaxUnsuccessfulPings int
}
