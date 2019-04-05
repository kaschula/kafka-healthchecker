package config

import "flag"

type AppConfig struct {
	FrequencySeconds int64
	KafkaAddress     string
	KafkaTopic       string
	HttpPort         string
}

func AppConfigFromFlags() AppConfig {
	address := flag.String("kafkaaddress", "localhost:9092", "address and port fo kafka server")
	defaultTopic := flag.String("topic", "health-check-test", "name of default topic")
	frequencySeconds := flag.Int64("freqseconds", 3, "Frequency of healthcheck in seconds")
	http := flag.String("httpport", "7070", "http port for health checker service")
	flag.Parse()

	return AppConfig{
		FrequencySeconds: *frequencySeconds,
		KafkaAddress:     *address,
		KafkaTopic:       *defaultTopic,
		HttpPort:         *http,
	}
}
