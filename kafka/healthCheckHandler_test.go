package kafka_test

import (
	"errors"
	"testing"
	"time"

	conf "github.com/kaschula/kafka-healthchecker/config"
	healthcheck "github.com/kaschula/kafka-healthchecker/healthcheck"
	k "github.com/kaschula/kafka-healthchecker/kafka"
)

type test struct {
	label      string
	expected   expect
	connection *connectionFake
	kafkaError error
}

func (t *test) getKafkaCLi() kafkaFake {
	return newConnectionFake(t.connection, t.kafkaError)
}

type expect struct {
	health    bool
	connected bool
	created   bool
	deleted   bool
	closed    bool
	msg       string
}

func TestTheHandler(t *testing.T) {
	config := conf.AppConfig{
		FrequencySeconds: 0,
		KafkaAddress:     "127.address",
		KafkaTopic:       "test-topic",
		HttpPort:         "n/a",
	}

	tests := []test{
		{
			"connection and cli return no errors, healthy status return",
			expect{health: true, connected: true, created: true, deleted: true, closed: true, msg: ""},
			newConnectionFakeNoError(nil, nil, nil),
			nil,
		},
		{
			"connection fails and connection methods never called",
			expect{health: false, connected: false, created: false, deleted: false, closed: false, msg: "Connection Error"},
			newConnectionFakeNoError(nil, nil, nil),
			errors.New("Connection Error"),
		},
		{
			"it connects but create fails",
			expect{health: false, connected: true, created: false, deleted: false, closed: false, msg: "Create Error"},
			newConnectionFakeNoError(errors.New("Create Error"), nil, nil),
			nil,
		},
		{
			"it connects, creates but delete fails",
			expect{health: false, connected: true, created: true, deleted: false, closed: false, msg: "Delete Error"},
			newConnectionFakeNoError(nil, errors.New("Delete Error"), nil),
			nil,
		},
		{
			"it connects, creates, delete but close fails",
			expect{health: false, connected: true, created: true, deleted: true, closed: false, msg: "Close Error"},
			newConnectionFakeNoError(nil, nil, errors.New("Close Error")),
			nil,
		},
	}

	for _, test := range tests {
		expected := test.expected
		connection := test.connection
		kafkaCli := test.getKafkaCLi()
		hc := healthcheck.New(false, time.Now(), "initialize")
		state := healthcheck.HealthCheckRepository{CurrentState: hc}

		handler := k.CreateHandler(config, &kafkaCli, &state)
		handler()
		assertResponseAndFake(t, state.GetState(), kafkaCli, *connection, expected, test.label)
	}
}

func assertResponseAndFake(
	t *testing.T,
	hc healthcheck.HealthCheck,
	kafkaCli kafkaFake,
	connection connectionFake,
	expected expect,
	label string,
) {
	if hc.Healthy != expected.health ||
		hc.Message != expected.msg ||
		kafkaCli.connected != expected.connected ||
		connection.created != expected.created ||
		connection.deleted != expected.deleted ||
		connection.closed != expected.closed {
		t.Log(label)
		t.Fatalf(
			"Expected : Actual. hc.Healthy = %v : %v, "+
				"connected = %v : %v, "+
				"created = %v : %v, "+
				"deleted= %v : %v, "+
				"closed= %v : %v, "+
				"error message %v : %v",
			expected.health,
			hc.Healthy,
			expected.connected,
			kafkaCli.connected,
			expected.created,
			connection.created,
			expected.deleted,
			connection.deleted,
			expected.closed,
			connection.closed,
			hc.Message,
			expected.msg,
		)
	}
}

type connectionFake struct {
	created       bool
	deleted       bool
	closed        bool
	createdReturn error
	deletedReturn error
	closedReturn  error
}

func (c *connectionFake) CreateTopics(name string, partitions, replication int) error {
	if c.createdReturn == nil {
		c.created = true
	}

	return c.createdReturn
}

func (c *connectionFake) DeleteTopics(name string) error {
	if c.deletedReturn == nil {
		c.deleted = true
	}

	return c.deletedReturn
}

func (c *connectionFake) Close() error {
	if c.closedReturn == nil {
		c.closed = true
	}

	return c.closedReturn
}

type kafkaFake struct {
	connected bool
	conn      k.KafkaConnection
	errReturn error
}

func (k *kafkaFake) Connect(network, address string) (k.KafkaConnection, error) {
	if k.errReturn != nil {
		return nil, k.errReturn
	}

	k.connected = true

	return k.conn, k.errReturn
}

func newConnectionFakeNoError(createError, deleteError, closedError error) *connectionFake {
	return &connectionFake{
		created:       false,
		deleted:       false,
		closed:        false,
		createdReturn: createError,
		deletedReturn: deleteError,
		closedReturn:  closedError,
	}
}

func newConnectionFake(connection *connectionFake, errReturn error) kafkaFake {
	return kafkaFake{connected: false, conn: connection, errReturn: errReturn}
}
