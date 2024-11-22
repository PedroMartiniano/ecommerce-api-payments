package messaging

import (
	"context"
	"encoding/json"

	"github.com/PedroMartiniano/ecommerce-api-payments/internal/application/services"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/configs/env"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/configs/logger"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/domain/dtos"
	amqp "github.com/rabbitmq/amqp091-go"
)

var log = logger.NewLogger()

type RabbitMqQueue struct {
	paymentService *services.PaymentService
}

type BodyResponse struct {
	OrderID        string  `json:"order_id"`
	UserID         string  `json:"user_id"`
	CardHolder     string  `json:"card_holder"`
	CardNumber     string  `json:"card_number"`
	ExpirationDate string  `json:"expiration_date"`
	CVV            string  `json:"cvv"`
	Amount         float64 `json:"amount"`
}

func NewRabbitMqQueue(paymentService *services.PaymentService) *RabbitMqQueue {
	return &RabbitMqQueue{
		paymentService: paymentService,
	}
}

func (mq *RabbitMqQueue) InitPaymentConsumer() error {
	queueURL := env.GetEnv("QUEUE_URL")
	conn, err := amqp.Dial(queueURL)
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
			var response BodyResponse

			err := json.Unmarshal(d.Body, &response)
			if err != nil {
				log.Errorf("Error unmarshalling message: %s", err.Error())
			}

			err = mq.paymentService.ProcessPaymentExecute(context.Background(), dtos.ProcessPaymentDTO{
				OrderID:        response.OrderID,
				UserID:         response.UserID,
				CardHolder:     response.CardHolder,
				CardNumber:     response.CardNumber,
				ExpirationDate: response.ExpirationDate,
				CVV:            response.CVV,
				Amount:         response.Amount,
			})
			if err != nil {
				log.Errorf("Error processing payment: %s", err.Error())
			}
		}
	}()

	log.Info("[*] Waiting for messages...")
	<-forever

	return nil
}
