package rabbitmq

import (
	"github.com/streadway/amqp"
)

// AMQP Object Structure
type AMQP struct {
	conn  *amqp.Connection
	Ch    *amqp.Channel
	Queue *amqp.Queue
}

// NewConnection : Creates New RabbitMQ Connection Object
func NewConnection(path string) (*AMQP, error) {
	conn, err := amqp.Dial(path)
	if err != nil {
		return nil, err
	}

	amqp := new(AMQP)
	amqp.conn = conn
	return amqp, nil
}

// CreateChannel : Creates New RabbitMQ Channel
func (c *AMQP) CreateChannel() error {
	ch, err := c.conn.Channel()
	c.Ch = ch

	if err != nil {
		return err
	}
	return nil
}

// DeclareQueue : Declares RabbitMQ Queue
func (c *AMQP) DeclareQueue() error {
	q, err := c.Ch.QueueDeclare(
		"message", // name can be taken from config
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	c.Queue = &q

	if err != nil {
		return err
	}

	return nil
}

// SendData : Sends Data
func (c *AMQP) SendData(data []byte) error {
	err := c.Ch.Publish(
		"",           // exchange
		c.Queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})
	if err != nil {
		return err
	}

	return nil
}

// RecieveData : Recieves Data
func (c *AMQP) RecieveData() (<-chan amqp.Delivery, error) {
	msgs, err := c.Ch.Consume(
		c.Queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		return msgs, err
	}

	return msgs, nil
}

// Close : Close connection
func (c *AMQP) Close() {
	c.conn.Close()
}
