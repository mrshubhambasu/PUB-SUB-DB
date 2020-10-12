package pub

import (
	"assignment/rabbitmq"
)

func InitService(rabbitMQURL string, data []byte) error {
	rabbitMQ, err := rabbitmq.NewConnection(rabbitMQURL)
	if err != nil {
		return err
	}
	defer rabbitMQ.Close()

	err = rabbitMQ.CreateChannel()
	if err != nil {
		return err
	}

	err = rabbitMQ.DeclareQueue()
	if err != nil {
		return err
	}

	// Sending File
	err = rabbitMQ.SendData(data)
	if err != nil {
		return err
	}
	return nil
}
