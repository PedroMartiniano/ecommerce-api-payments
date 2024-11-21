package services

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-payments/internal/application/ports/gateways"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/application/ports/repositories"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/domain/dtos"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/domain/entities"
)

type PaymentService struct {
	paymentRepository repositories.IPaymentRepository
	paymentGateway    gateways.IPaymentGateway
}

func NewPaymentService(paymentRepository repositories.IPaymentRepository, paymentGateway gateways.IPaymentGateway) *PaymentService {
	return &PaymentService{
		paymentRepository: paymentRepository,
		paymentGateway:    paymentGateway,
	}
}

func (p *PaymentService) ProcessPaymentExecute(c context.Context, dto dtos.ProcessPaymentDTO) error {
	err := p.paymentGateway.ProcessPayment(c, dto)
	if err != nil {
		return err
	}

	paymentEnt := entities.NewPaymentEntity(
		dto.OrderID,
		dto.UserID,
		"Credit Card",
		"Success",
		dto.Amount,
	)

	return p.paymentRepository.SavePayment(c, paymentEnt)
}
