package kafka

import (
	// "fmt"

	k "github.com/segmentio/kafka-go"
)

// Will have to make a concreate adapater to test the segmentio package.
// To do this run Docker with kafka and test the module against that

// kafka-go interfaces
// Wont need these

// type KafkaGoDials interface {
// 	Dial(network, address string) (KafkaGoConn, error)
// }

// type KafkaGoConn interface {
// 	CreateTopics(topics ...k.TopicConfig) error
// 	DeleteTopics(topics ...string) error
// 	Close() error
// }

func NewSegmentionAdapaterDialer() *segmentioGoAdapaterDialer {
	dialer := k.DefaultDialer

	return &segmentioGoAdapaterDialer{dialer: dialer}
}

type segmentioGoAdapaterDialer struct {
	dialer *k.Dialer
}

func (kafkaGo *segmentioGoAdapaterDialer) Connect(network, address string) (KafkaConnection, error) {
	conn, err := kafkaGo.dialer.Dial(network, address)

	return &segmentioGoAdapaterConn{conn}, err
}

type segmentioGoAdapaterConn struct {
	conn *k.Conn
}

func (s *segmentioGoAdapaterConn) CreateTopics(topic string, partitions, replicationFactor int) error {
	config := k.TopicConfig{Topic: topic, NumPartitions: partitions, ReplicationFactor: replicationFactor}

	return s.conn.CreateTopics(config)
}

func (s *segmentioGoAdapaterConn) DeleteTopics(topic string) error {
	return s.conn.DeleteTopics(topic)
}

func (s *segmentioGoAdapaterConn) Close() error {
	return s.conn.Close()
}
