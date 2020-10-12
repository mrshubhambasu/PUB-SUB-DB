package main

import (
	"assignment/config"
	"assignment/pub"
	"assignment/sub"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

func main() {
	// Setting Logger Format
	log.SetFormatter(&log.JSONFormatter{})

	// Loading Configurations
	config.Load()

	// Reading RabbitMQ Path From Config
	rabbitMQURL := config.ReadEnvString("RABBITMQ_URL")

	// Reading DSN of Database From Config
	dsn := config.ReadEnvString("DSN")

	file, err := ioutil.ReadFile("inputfile.json")
	if err != nil {
		log.Error(err)
	}

	// Calling Publisher
	err = pub.InitService(rabbitMQURL, file)
	if err != nil {
		log.Error(err)
	}

	// Calling Subscriber
	err = sub.InitService(rabbitMQURL, dsn)
	if err != nil {
		log.Error(err)
	}
}
