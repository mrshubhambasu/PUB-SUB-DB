package main

import (
	"assignment/config"
	"assignment/rabbitmq"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	// Initializig DB Connection
	dsn := config.ReadEnvString("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("err", err)
	}

	// Get Message Channel
	msgs := conn.RecieveData()
	forever := make(chan bool)

	// Waiting For Messages
	go func() {
		for d := range msgs {
			log.Info("Recieved Data")
			go pushToDatabase(d.Body, db)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
