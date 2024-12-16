package rabbitmq

import "github.com/streadway/amqp"

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func Connect(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		Conn:    conn,
		Channel: ch,
	}, nil
}

//// Publish отправляет сообщение в очередь
//func Publish(conn *amqp.Connection, queueName string, message []byte) error {
//	ch, err := conn.Channel()
//	if err != nil {
//		return err
//	}
//	defer ch.Close()
//
//	q, err := ch.QueueDeclare(
//		queueName, // name
//		false,     // durable
//		false,     // delete when unused
//		false,     // exclusive
//		false,     // no-wait
//		nil,       // arguments
//	)
//	if err != nil {
//		return err
//	}
//
//	err = ch.Publish(
//		"",     // exchange
//		q.Name, // routing key
//		false,  // mandatory
//		false,  // immediate
//		amqp.Publishing{
//			ContentType: "application/json",
//			Body:        message,
//		})
//	return err
//}

func (r *RabbitMQ) Close() error {
	if r.Channel != nil {
		r.Channel.Close()
	}
	if r.Conn != nil {
		return r.Conn.Close()
	}
	return nil
}
