package rabbitmq

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

// AMQP Object Structure
type AMQP struct {
	conn  *amqp.Connection
	Ch    *amqp.Channel
	Queue *amqp.Queue
}

// NewConnection : Creates New RabbitMQ Connection Object
func NewConnection(path string) *AMQP {
	conn, err := amqp.Dial(path)
	failOnError(err, "Failed to connect to RabbitMQ")

	amqp := new(AMQP)
	amqp.conn = conn
	return amqp
}

// CreateChannel : Creates New RabbitMQ Channel
func (c *AMQP) CreateChannel() {
	ch, err := c.conn.Channel()
	c.Ch = ch

	failOnError(err, "Failed to open a channel")
}

// DeclareQueue : Declares RabbitMQ Queue
func (c *AMQP) DeclareQueue(name string) {
	q, err := c.Ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	c.Queue = &q
	failOnError(err, "Failed to declare a queue")
}

// SendData : Sends Data
func (c *AMQP) SendData(data []byte) {
	err := c.Ch.Publish(
		"",           // exchange
		c.Queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})
	failOnError(err, "Failed to publish a message")
}

// RecieveData : Recieves Data
func (c *AMQP) RecieveData() <-chan amqp.Delivery {
	msgs, err := c.Ch.Consume(
		c.Queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	failOnError(err, "Failed to register a consumer")

	return msgs
}

// Handles Errors Gracefully
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err) // not logrus
	}
}
