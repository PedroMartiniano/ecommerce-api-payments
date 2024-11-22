package queue

import (
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/infra/adapters"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/infra/messaging"
)

func InitQueueConsumer() error {
	paymentService := adapters.NewPaymentServiceAdapter(configs.DB)
	queue := messaging.NewRabbitMqQueue(paymentService)

	err := queue.InitPaymentConsumer()
	return err
}
