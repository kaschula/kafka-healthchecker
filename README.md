# Kafka-HealthChecker

A health checking service that monitors kafka.

## Motivation
This project was created for a friend who required an easy way of of checking the current status of a kafka set up.

## How it works

This service provides a web endpoint to retrieve the latest healthckeck status.

Based on configuration the application runs a health check at a defined interval in the background. 

This health check consists of: 
- Connecting to Kafka
- Creating a dummy Topic
- Deleting the dummy topic
- Closing the connection

If all these tasks are preformed successfully then a kafka service is considered healthy.

If any task fails then service health check will set the status no false (unhealthy) and return any error messages provided by the kafka client. 

The `/check` path will always return the latest update.

## Installation
#### Prerequisites

Golang 1.10.3 or higher

#### Install

run: 

```go get github.com/kaschula/kafka-healthchecker```

```go install```

## How to use?

Once installed and assuming you $GOPATH/bin directory is added to your $PATH. From the command line run:

```kafka-healthchecker```

Once service is running visit `localhost:8080/check`url to get the latest health check information. This returns a json response.

The service config can be adjust by passing in flags to the command call. This illustrates setting a value for each config item.  

```kafka-healthchecker --kafkaaddress="127.0.0.1:9099" --topic=health-check-topic-name --freqseconds=5 --httpport=8080```

Config items have default values so only options that you want to set need to be passed in on the call.

### Example

Following [this guide](https://kafka.apache.org/quickstart) to get started with kafka.

Once zookeeper and kafka are up running start build and kafka-healthchecker service. The default port for kafka based on the default kafka config is `127.0.0.1:9092`

### Config and Defaults
Table of the availible config and the default values

| Item | flag | default | description |
|------|:-----|:--------|:------------|
| Kafka Address | --kafkaaddress | "localhost:9092" | address and port for the kafka service server |
| Health Check Topic | --topic | "health-check-test" | Name of default topic |
| Frequency of health check | --freqseconds | 3 | "Frequency at which the health check is run in seconds" |
| Health checker service port | --httpport | "7070" | The HTTP port for health checker service |


## API Reference

| Method | uri | response status code | type | body |
|--------|-----|----------------------|------|------|
| GET | `/check` | 200| json | ```{healthy: boolean, time: rfc3339 timestamp, message: string, contains status information ```
## Screenshots
Include logo/demo screenshot etc.

## Tech + frameworks used

The Golang language core.

Apache Kafka

### Built with:

[segmentio/kafka-go](https://github.com/segmentio/kafka-go)

## Features
What makes your project stand out?

## Tests

The testing is done using the standard `go test`.

To run the tests from the root of the project run:

```go test ./...```

Currently there is no integration tests. This is still in development in a seperate branch.


## Contribute

Anyone is welcome to contribute please send pull requests. 

## Credits
- Kafka Client [segmentio/kafka-go](https://github.com/segmentio/kafka-go)
- This [article](https://medium.com/@meakaakka/a-beginners-guide-to-writing-a-kickass-readme-7ac01da88ab3) for a guide to writing this README

## To Do:
- Add Integration tests for the kafka-go library adapter working with docker
- Add sending a messages to a topic to the health check process
- Increase test coverage
- Set up build pipeline 
- Test running on linux environment 

## License
Licenses under the Apache License 2.0
