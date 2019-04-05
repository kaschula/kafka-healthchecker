package kafka

import (
	"time"

	"github.com/kaschula/kafka-healthchecker/config"
	"github.com/kaschula/kafka-healthchecker/healthcheck"
)

func CreateHandler(config config.AppConfig, kafka KafkaCli, currentHealth *healthcheck.HealthCheckRepository) func() {
	return func() {
		conn, err := kafka.Connect("tcp", config.KafkaAddress)
		if err != nil {
			currentHealth.SetState(newFailedHealthCheck(err.Error()))

			return
		}

		err = conn.CreateTopics("string", 1, 1)
		if err != nil {
			currentHealth.SetState(newFailedHealthCheck(err.Error()))

			return
		}

		err = conn.DeleteTopics("string")
		if err != nil {
			currentHealth.SetState(newFailedHealthCheck(err.Error()))

			return
		}

		err = conn.Close()
		if err != nil {
			currentHealth.SetState(newFailedHealthCheck(err.Error()))

			return
		}

		hc := healthcheck.New(true, time.Now(), "")
		currentHealth.SetState(hc)
	}
}

func newFailedHealthCheck(message string) healthcheck.HealthCheck {
	return healthcheck.New(false, time.Now(), message)
}
