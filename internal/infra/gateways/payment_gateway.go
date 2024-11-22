package gateways

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-payments/internal/application/ports/gateways"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/configs/logger"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/domain/dtos"
)

var log = logger.NewLogger()

type StripePaymentGateway struct{}

func NewStripePaymentGateway() gateways.IPaymentGateway {
	return &StripePaymentGateway{}
}

func (s StripePaymentGateway) ProcessPayment(c context.Context, dto dtos.ProcessPaymentDTO) error {
	log.Info("Processing payment...")
	return nil
}
