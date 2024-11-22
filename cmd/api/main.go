package main

import (
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/infra/queue"
)

func main() {
	configs.InitConfig()

	err := queue.InitQueueConsumer()
	panic(err)
}
