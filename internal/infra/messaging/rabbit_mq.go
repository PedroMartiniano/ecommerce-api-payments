package messaging

import (
	"log"

	"github.com/PedroMartiniano/ecommerce-api-payments/internal/application/services"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/configs"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMqQueue struct {
	paymentService *services.PaymentService
}

func NewRabbitMqQueue(paymentService *services.PaymentService) *RabbitMqQueue {
	return &RabbitMqQueue{
		paymentService: paymentService,
	}
}

func (mq *RabbitMqQueue) InitPaymentConsumer() error {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"process-payment",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			configs.Logger.Infof(" [x] Received %s", d.Body)
		}
	}()

	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}
