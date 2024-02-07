package rmq

import (
	"errors"
	"github.com/streadway/amqp"
)

var creationErr = errors.New("connect to radditmq")

type consumer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
	messages   <-chan amqp.Delivery
}

func newConsumer(addr string, queueName string) (consumer, error) {
	conn, err := amqp.Dial(addr)
	if err != nil {
		return consumer{}, errors.Join(creationErr, err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return consumer{}, errors.Join(creationErr, err)
	}

	queue, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return consumer{}, errors.Join(creationErr, err)
	}

	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return consumer{}, errors.Join(creationErr, err)
	}

	return consumer{
		connection: conn,
		channel:    ch,
		queue:      queue,
		messages:   msgs,
	}, nil
}

func (c *consumer) close() {
	_ = c.channel.Close()
	_ = c.connection.Close()
}
