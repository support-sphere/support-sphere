package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
	"github.com/support-sphere/support-sphere/internal/core/domain"
)

type TicketPublisher struct {
	conn *amqp.Connection
}

func NewTicketPublisher(conn *amqp.Connection) *TicketPublisher {
	return &TicketPublisher{conn: conn}
}

func (p *TicketPublisher) PublishTicket(ticket domain.Ticket) error {
	ch, err := p.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"ticket_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	body, err := json.Marshal(ticket)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	log.Printf(" [x] Sent %s", body)
	return nil
}
