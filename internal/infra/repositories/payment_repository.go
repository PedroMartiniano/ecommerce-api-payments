package repositories

import (
	"context"
	"database/sql"

	"github.com/PedroMartiniano/ecommerce-api-payments/internal/application/ports/repositories"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/configs/errors"
	"github.com/PedroMartiniano/ecommerce-api-payments/internal/domain/entities"
)

type PaymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) repositories.IPaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

func (p PaymentRepository) SavePayment(c context.Context, ent entities.Payments) error {
	query := `
		INSERT INTO payments(id, order_id, customer_id, total_amount, payment_method, payment_date, payment_status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	stmt, err := p.db.Prepare(query)
	if err != nil {
		return errors.NewError(errors.ErrInternalServer, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		c,
		ent.ID,
		ent.OrderID,
		ent.CustomerID,
		ent.TotalAmount,
		ent.PaymentMethod,
		ent.PaymentDate,
		ent.PaymentStatus,
		ent.CreatedAt,
	)
	if err != nil {
		return errors.NewError(errors.ErrInternalServer, err)
	}

	return nil
}
