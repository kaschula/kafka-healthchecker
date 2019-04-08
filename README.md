# Kafka-HealthChecker

A health checking service that monitors kafka.

## Motivation
This project was created for a friend who required an easy way of checking the current status of a kafka set up.

## How it works

This service provides a web endpoint to retrieve the latest health ckeck status.

Based on configuration the application runs a health check at a defined interval in the background by polling kafka. 

This health check consists of: 
- Connecting to Kafka
- Creating a Topic
- Deleting the new topic
- Closing the connection

If all these tasks are performed successfully then a kafka service is considered healthy.

If any task fails then the health check's `healthy` status will be set to false (unhealthy) and return any error messages provided by the kafka client. 

The `/check` path will return the latest update.

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

Once service is running visit `localhost:8080/check` URL to get the latest health check information. This returns a JSON response.

The service config can be adjusted by passing in arguments to the command call. The example below illustrates setting a value for each config item.  

```kafka-healthchecker --kafkaaddress="127.0.0.1:9099" --topic=health-check-topic-name --freqseconds=5 --httpport=8080```

Config items have default values. Only options that you want to set need to be passed in on the call.

### Config and Defaults
Table of the available config and the default values

| Item | argument | default | description |
|------|:-----|:--------|:------------|
| Kafka Address | --kafkaaddress | "localhost:9092" | address and port for the kafka service server |
| Health Check Topic | --topic | "health-check-test" | Name of default topic |
| Frequency of health check | --freqseconds | 3 | "Frequency at which the health check is run in seconds" |
| Health checker service port | --httpport | "7070" | The HTTP port for health checker service |

### Example

Following [this guide](https://kafka.apache.org/quickstart) to get started with kafka.

Once zookeeper and kafka are up running, build and start the kafka-healthchecker service. The default port for kafka based on the default kafka config in the Apache documentation is `127.0.0.1:9092`

## API Reference

| Method | uri | response status code | type | body |
|--------|-----|----------------------|------|------|
| GET | `/check` | 200| json | ```{healthy: boolean, time: rfc3339 timestamp, message: string, contains status information ```


## Tech + Frameworks used

The Golang language core.

Apache Kafka

### Built with:

[segmentio/kafka-go](https://github.com/segmentio/kafka-go)

## Tests

The testing is done using the standard `go test`.

To run the tests, from the root of the project run:

```go test ./...```

Currently there is no integration tests. This is still in development in a seperate branch.


## Contribute

Anyone is welcome to contribute please send pull requests. Pull request should out line the problem found and the solution developed to solve.

## Credits
- Kafka Client [segmentio/kafka-go](https://github.com/segmentio/kafka-go)
- This [article](https://medium.com/@meakaakka/a-beginners-guide-to-writing-a-kickass-readme-7ac01da88ab3) for a guide to writing this README

## To Do:
- Add Integration tests for the kafka-go library adapter working with docker
- Add sending a message to a topic to the health check process
- Increase test coverage
- Set up build pipeline 
- Test running on Linux environment 

## License
Licenses under the Apache License 2.0
