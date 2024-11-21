package adapters

import (
	"database/sql"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/application/services"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/infra/gateways"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/infra/repositories"
)

func NewPaymentServiceAdapter(db *sql.DB) *services.PaymentService {
	paymentGateway := gateways.NewStripePaymentGateway()
	paymentRepository := repositories.NewPaymentRepository(db)

	paymentService := services.NewPaymentService(paymentRepository, paymentGateway)
	return paymentService
}
