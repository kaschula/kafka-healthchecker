package healthcheck_test

import (
	"fmt"
	"testing"
	"time"

	hc "github.com/kaschula/kafka-healthchecker/healthcheck"
)

func TestItCreatesANewHealthCheckObject(t *testing.T) {
	tests := []struct {
		health  bool
		message string
		time    time.Time
	}{
		{true, "msg1", time.Now()},
		{false, "msg2", time.Now()},
	}

	for _, test := range tests {
		healthCheck := hc.New(test.health, test.time, test.message)

		if healthCheck.Healthy != test.health ||
			healthCheck.Time != test.time ||
			healthCheck.Message != test.message {
			t.Fatalf("Error %v does not equal %v", healthCheck, test)

			return
		}

		return
	}
}

func TestItTheHealthCheckCanMarshallToJson(t *testing.T) {
	checkTime := time.Now()
	timeString := checkTime.Format(time.RFC3339)

	tests := []struct {
		health   bool
		time     time.Time
		message  string
		expected []byte
	}{
		{true, checkTime, "msg1", createJsonString("true", timeString, "msg1")},
		{false, checkTime, "msg2", createJsonString("false", timeString, "msg2")},
	}

	for _, test := range tests {
		healthCheck, err := hc.New(test.health, test.time, test.message).ToJson()

		if err != nil {
			t.Fatalf("Error: %v", err)

			return
		}

		if string(healthCheck) != string(test.expected) {
			t.Fatalf("Error %v does not equal %v", string(healthCheck), test.expected)

			return
		}

		return
	}
}

func createJsonString(health, time, message string) []byte {
	return []byte(fmt.Sprintf("{\"healthy\":\"%v\",\"time\":\"%v\",\"message\":\"%v\"}", health, time, message))
}
