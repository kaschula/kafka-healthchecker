package scheduler_test

import (
	"fmt"
	"testing"

	config "github.com/kaschula/kafka-healthchecker/config"
	scheduler "github.com/kaschula/kafka-healthchecker/scheduler"
)

func TestThatAHandlerIsRunCalled(t *testing.T) {
	appConfig := config.AppConfig{FrequencySeconds: int64(1)}
	schedule := scheduler.New(appConfig)

	expected := "value"
	channelString := make(chan string)

	handler := func() {
		channelString <- expected
	}

	go schedule.Run(handler)

	received := <-channelString
	fmt.Println("Received")
	if received != expected {
		t.Fatalf("Expected '%v' to be '%v'", received, expected)
	}
}
