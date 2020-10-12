package sub

import (
	"assignment/datastore"
	"assignment/model"
	"assignment/rabbitmq"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func InitService(rabbitMQURL, dsn string) error {
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

	db, err := datastore.Init(dsn)
	if err != nil {
		return err
	}

	// Get Message Channel
	msgs, err := rabbitMQ.RecieveData()
	if err != nil {
		return err
	}

	forever := make(chan bool)

	dataModel := model.Data{}

	// Waiting For Messages
	go func() {
		for d := range msgs {
			dataModel, _ := byteToStruct(d.Body, dataModel)
			datastore.Push(dataModel, db)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
	return nil
}

// This function populates the struct with data from byte array
func byteToStruct(bytes []byte, dataModel model.Data) (model.Data, error) {
	err := json.Unmarshal(bytes, &dataModel)
	if err != nil {
		return model.Data{}, err
	}
	return dataModel, nil
}
