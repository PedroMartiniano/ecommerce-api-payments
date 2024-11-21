package entities

import (
	"time"

	"github.com/google/uuid"
)

type Payments struct {
	ID            string    `json:"id"`
	OrderID       string    `json:"order_id"`
	CustomerID    string    `json:"customer_id"`
	TotalAmount   float64    `json:"total_amount"`
	PaymentMethod string    `json:"payment_method"`
	PaymentDate   time.Time    `json:"payment_date"`
	PaymentStatus string    `json:"payment_status"`
	CreatedAt     time.Time `json:"created_at"`
}

func NewPaymentEntity(orderID, customerID, paymentMethod, paymentStatus string, totalAmount float64) Payments {
	uuid, _ := uuid.NewV7()


	return Payments{
		ID:            uuid.String(),
		OrderID:       orderID,
		CustomerID:    customerID,
		TotalAmount:   totalAmount,
		PaymentMethod: paymentMethod,
		PaymentDate:   time.Now(),
		PaymentStatus: paymentStatus,
		CreatedAt:     time.Now(),
	}
}
