package rmq

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
)

type producer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
}

func newProducer(addr string, queueName string) (producer, error) {
	conn, err := amqp.Dial(addr)
	if err != nil {
		return producer{}, errors.Join(creationErr, err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return producer{}, errors.Join(creationErr, err)
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
		return producer{}, errors.Join(creationErr, err)
	}

	return producer{
		connection: conn,
		channel:    ch,
		queue:      queue,
	}, nil
}

func (p *producer) publish(id string, resp jobResponse) error {
	body, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("response marshalling to json: %w", err)
	}
	return p.channel.Publish(
		"",
		p.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			MessageId:   id,
			Body:        body,
		},
	)
}

func (p *producer) close() {
	_ = p.channel.Close()
	_ = p.connection.Close()
}
