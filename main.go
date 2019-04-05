package main

import (
	"fmt"
	"net/http"
	"time"

	conf "github.com/kaschula/kafka-healthchecker/config"
	"github.com/kaschula/kafka-healthchecker/healthcheck"
	kafkaService "github.com/kaschula/kafka-healthchecker/kafka"
	schedulerService "github.com/kaschula/kafka-healthchecker/scheduler"
)

func main() {
	config := conf.AppConfigFromFlags()
	fmt.Println("config", config)
	scheduler := schedulerService.New(config)
	kafkaCli := createKafkaCli(config)

	state := healthcheck.NewRepository(healthcheck.New(false, time.Now(), "Initializing"))

	// Setup server
	http.HandleFunc("/check", func(res http.ResponseWriter, request *http.Request) {
		response, err := state.GetState().ToJson()

		if err != nil {
			response = []byte(fmt.Sprintf("{\"error\":\"%v\"}", err.Error()))
		}

		res.Write(response)
	})

	fmt.Println("Starting scheduler....")
	go scheduler.Run(kafkaService.CreateHandler(config, kafkaCli, &state))

	fmt.Println("Starting Server....")
	http.ListenAndServe(":"+config.HttpPort, nil)
}

func createKafkaCli(config conf.AppConfig) kafkaService.KafkaCli {
	return kafkaService.NewSegmentionAdapaterDialer()
}
