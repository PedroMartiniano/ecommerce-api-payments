package gateways

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-payments/internal/application/ports/gateways"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/domain/dtos"
)

type StripePaymentGateway struct{}

func NewStripePaymentGateway() gateways.IPaymentGateway {
	return &StripePaymentGateway{}
}

func (s StripePaymentGateway) ProcessPayment(c context.Context, dto dtos.ProcessPaymentDTO) error {
	return nil
}
