package kafka

type KafkaCli interface {
	Connect(network, address string) (KafkaConnection, error)
}

type KafkaConnection interface {
	// This should topic, singular
	CreateTopics(name string, partitions, replication int) error
	DeleteTopics(name string) error
	Close() error
}
