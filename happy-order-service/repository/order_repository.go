package repository

import (
	"context"
	"database/sql"
	"happy-order-service/model/domain"
)

type OrderRepository interface {
	Save(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order
	UpdateStatusPayment(ctx context.Context, tx *sql.Tx, orderId string)
	UpdateFeedbackDone(ctx context.Context, tx *sql.Tx, orderId string)
	Delete(ctx context.Context, tx *sql.Tx, orderId string)
	FindByIdOrder(ctx context.Context, tx *sql.Tx, orderId string) (domain.Order, error)
	FindByIdUserAndIdOrder(ctx context.Context, tx *sql.Tx, userId uint, orderId string) (domain.Order, error)
	FindAllByIdUser(ctx context.Context, tx *sql.Tx, userId uint) []domain.Order
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Order
}
