package scheduler

import (
	"time"

	app "github.com/kaschula/kafka-healthchecker/config"
)

func New(config app.AppConfig) Scheduler {
	return Scheduler{frequencySeconds: config.FrequencySeconds}
}

type Scheduler struct {
	frequencySeconds int64
}

func (s Scheduler) Run(handler func()) {
	for {
		time.Sleep(time.Duration(s.time()))
		handler()
	}
}

func (s Scheduler) time() int64 {
	return (1000 * s.frequencySeconds) * int64(time.Millisecond)
}
