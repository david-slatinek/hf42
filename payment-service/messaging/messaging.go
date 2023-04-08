package messaging

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

type Messaging struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   *amqp.Queue
}

func NewMessaging() (*Messaging, error) {
	msg := &Messaging{}
	err := msg.init()
	return msg, err
}

func (receiver *Messaging) init() error {
	conn, err := amqp.Dial(os.Getenv("AMQP_URL"))
	if err != nil {
		return err
	}
	receiver.conn = conn

	ch, err := receiver.conn.Channel()
	if err != nil {
		return err
	}
	receiver.channel = ch

	q, err := ch.QueueDeclare(
		os.Getenv("EXCHANGE_QUEUE_NAME"),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	receiver.queue = &q
	return nil
}

func (receiver *Messaging) Close() {
	if receiver.conn != nil && !receiver.conn.IsClosed() {
		err := receiver.conn.Close()
		if err != nil {
			log.Printf("conn close error: %v", err)
		}
	}
	if receiver.channel != nil && !receiver.channel.IsClosed() {
		err := receiver.channel.Close()
		if err != nil {
			log.Printf("channel close error: %v", err)
		}
	}
}

func (receiver *Messaging) Consume() (<-chan amqp.Delivery, error) {
	return receiver.channel.Consume(receiver.queue.Name, "", true, false, false, false, nil)
}
