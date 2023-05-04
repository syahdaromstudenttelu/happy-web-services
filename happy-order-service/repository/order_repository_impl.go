package repository

import (
	"context"
	"database/sql"
	"errors"
	"happy-order-service/helper"
	"happy-order-service/lib"
	"happy-order-service/model/domain"
	"time"
)

type OrderRepositoryImpl struct{}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (repository *OrderRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order {
	sqlExecOne := "INSERT INTO order_product(order_id, total_price) VALUES (?, ?)"
	totalPriceOrder := order.Quantity * order.Price
	orderIdGenerated := lib.GetRandomStdId(48)
	_, err := tx.ExecContext(ctx, sqlExecOne, orderIdGenerated, totalPriceOrder)
	helper.DoPanicIfError(err)

	sqlExecTwo := "INSERT INTO order_detail(id_user, id_product, id_order, price, quantity) VALUES (?, ?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, sqlExecTwo, order.IdUser, order.IdProduct, orderIdGenerated, order.Price, order.Quantity)
	helper.DoPanicIfError(err)

	order.IdOrder = orderIdGenerated
	order.TotalPrice = totalPriceOrder
	order.OrderedDate = time.Now()
	order.ExpiredDate = time.Now().Add(time.Hour * 5)
	order.StatusPayment = false

	return order
}

func (repository *OrderRepositoryImpl) UpdateStatusPayment(ctx context.Context, tx *sql.Tx, orderId string) {
	sqlExec := "UPDATE order_detail SET status_payment = true WHERE id_order = ?"
	_, err := tx.ExecContext(ctx, sqlExec, orderId)
	helper.DoPanicIfError(err)
}

func (repository *OrderRepositoryImpl) UpdateFeedbackDone(ctx context.Context, tx *sql.Tx, orderId string) {
	sqlExec := "UPDATE order_detail SET feedback_done = true WHERE id_order = ?"
	_, err := tx.ExecContext(ctx, sqlExec, orderId)
	helper.DoPanicIfError(err)
}

func (repository *OrderRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, orderId string) {
	sqlExec := "DELETE FROM order_product WHERE order_id = ?"
	_, err := tx.ExecContext(ctx, sqlExec, orderId)
	helper.DoPanicIfError(err)
}

func (repository *OrderRepositoryImpl) FindByIdOrder(ctx context.Context, tx *sql.Tx, orderId string) (domain.Order, error) {
	sqlQuery := "SELECT od.id_user as id_user, od.id_product as id_product, op.order_id as order_id, od.price as price, od.quantity as quantity, op.total_price as total_price, op.ordered_date as ordered_date, ADDTIME(op.ordered_date, '5:0:0') as expired_date, od.status_payment as status_payment, od.feedback_done as feedback_done FROM order_detail AS od JOIN order_product as op ON (od.id_order = op.order_id) WHERE op.order_id = ?"
	rows, err := tx.QueryContext(ctx, sqlQuery, orderId)
	helper.DoPanicIfError(err)
	defer rows.Close()

	order := domain.Order{}
	if rows.Next() {
		err := rows.Scan(&order.IdUser, &order.IdProduct, &order.IdOrder, &order.Price, &order.Quantity, &order.TotalPrice, &order.OrderedDate, &order.ExpiredDate, &order.StatusPayment, &order.FeedbackDone)
		helper.DoPanicIfError(err)

		return order, nil
	} else {
		return order, errors.New("order is not found")
	}
}

func (repository *OrderRepositoryImpl) FindByIdUserAndIdOrder(ctx context.Context, tx *sql.Tx, userId uint, orderId string) (domain.Order, error) {
	sqlQuery := "SELECT od.id_user as id_user, od.id_product as id_product, op.order_id as order_id, od.price as price, od.quantity as quantity, op.total_price as total_price, op.ordered_date as ordered_date, ADDTIME(op.ordered_date, '5:0:0') as expired_date, od.status_payment as status_payment, od.feedback_done as feedback_done FROM order_detail AS od JOIN order_product as op ON (od.id_order = op.order_id) WHERE od.id_user = ? AND op.order_id = ?"
	rows, err := tx.QueryContext(ctx, sqlQuery, userId, orderId)
	helper.DoPanicIfError(err)
	defer rows.Close()

	order := domain.Order{}
	if rows.Next() {
		err := rows.Scan(&order.IdUser, &order.IdProduct, &order.IdOrder, &order.Price, &order.Quantity, &order.TotalPrice, &order.OrderedDate, &order.ExpiredDate, &order.StatusPayment, &order.FeedbackDone)
		helper.DoPanicIfError(err)

		return order, nil
	} else {
		return order, errors.New("order is not found")
	}
}

func (repository *OrderRepositoryImpl) FindAllByIdUser(ctx context.Context, tx *sql.Tx, userId uint) []domain.Order {
	sqlQuery := "SELECT od.id_user as id_user, od.id_product as id_product, op.order_id as order_id, od.price as price, od.quantity as quantity, op.total_price as total_price, op.ordered_date as ordered_date, ADDTIME(op.ordered_date, '5:0:0') as expired_date, od.status_payment as status_payment, od.feedback_done as feedback_done FROM order_detail AS od JOIN order_product as op ON (od.id_order = op.order_id) WHERE id_user = ?"
	rows, err := tx.QueryContext(ctx, sqlQuery, userId)
	helper.DoPanicIfError(err)
	defer rows.Close()

	orders := []domain.Order{}
	for rows.Next() {
		order := domain.Order{}
		err := rows.Scan(&order.IdUser, &order.IdProduct, &order.IdOrder, &order.Price, &order.Quantity, &order.TotalPrice, &order.OrderedDate, &order.ExpiredDate, &order.StatusPayment, &order.FeedbackDone)
		helper.DoPanicIfError(err)
		orders = append(orders, order)
	}

	return orders
}

func (repository *OrderRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Order {
	sqlQuery := "SELECT od.id_user as id_user, od.id_product as id_product, op.order_id as order_id, od.price as price, od.quantity as quantity, op.total_price as total_price, op.ordered_date as ordered_date, ADDTIME(op.ordered_date, '5:0:0') as expired_date, od.status_payment as status_payment, od.feedback_done as feedback_done FROM order_detail AS od JOIN order_product as op ON (od.id_order = op.order_id)"
	rows, err := tx.QueryContext(ctx, sqlQuery)
	helper.DoPanicIfError(err)
	defer rows.Close()

	orders := []domain.Order{}
	for rows.Next() {
		order := domain.Order{}
		err := rows.Scan(&order.IdUser, &order.IdProduct, &order.IdOrder, &order.Price, &order.Quantity, &order.TotalPrice, &order.OrderedDate, &order.ExpiredDate, &order.StatusPayment, &order.FeedbackDone)
		helper.DoPanicIfError(err)
		orders = append(orders, order)
	}

	return orders
}
