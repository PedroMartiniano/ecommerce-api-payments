package main

import (
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/infra/adapters"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/infra/messaging"
)

func main() {
	configs.InitConfig()
	paymentService := adapters.NewPaymentServiceAdapter(configs.DB)
	queue := messaging.NewRabbitMqQueue(paymentService)

	err := queue.InitPaymentConsumer()
	panic(err)
}
