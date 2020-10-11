package main

import (
	"assignment/config"
	"assignment/rabbitmq"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

func main() {
	// Setting Logger Format
	log.SetFormatter(&log.JSONFormatter{})

	// Loading Configurations
	config.Load()

	// Setting Up rabbitMQ
	conn := rabbitmq.NewConnection(config.ReadEnvString("RABBITMQ_URL"))
	conn.CreateChannel()
	conn.DeclareQueue(config.ReadEnvString("QUEUE_NAME"))

	//TBD: Close connection in defer

	// Reading JSON file to send
	body, err := ioutil.ReadFile("inputfile.json")
	if err != nil {
		log.Error(err)
	}

	// Sending File
	conn.SendData(body)
}
