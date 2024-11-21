package gateways

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-payments/internal/domain/dtos"
)

type IPaymentGateway interface {
	ProcessPayment(context.Context, dtos.ProcessPaymentDTO) error
}
