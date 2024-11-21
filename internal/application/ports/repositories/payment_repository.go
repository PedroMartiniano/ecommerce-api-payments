package repositories

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-payments/internal/domain/entities"
)

type IPaymentRepository interface {
	SavePayment(context.Context, entities.Payments) error
}
