package test

import (
	"context"
	"happy-order-service/app"
	"happy-order-service/helper"
	"happy-order-service/model/domain"
	"happy-order-service/util"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestTemp(t *testing.T) {
	config, err := util.LoadConfig("../")
	helper.DoPanicIfError(err)

	db := app.NewDB(config.DBDriver, config.DBSource)
	ctx := context.Background()

	tx, err := db.Begin()
	helper.DoPanicIfError(err)

	sqlQuery := "SELECT od.id_user as id_user, od.id_product as id_product, op.order_id as order_id, od.price as price, od.quantity as quantity, op.total_price as total_price, op.ordered_date as ordered_date, ADDTIME(op.ordered_date, '5:0:0') as expired_date, od.status_payment as status_payment FROM order_detail AS od JOIN order_product as op ON (od.id_order = op.order_id) WHERE od.id_user = ? AND op.order_id = ?"
	rows, err := tx.QueryContext(ctx, sqlQuery, 1, "Bn-P_ZOvX5hN307B6qvPSftz-4y1B5Z6SKwAZrjuMdiJTgM3")
	helper.DoPanicIfError(err)
	defer rows.Close()

	order := domain.Order{}
	if rows.Next() {
		err := rows.Scan(&order.IdUser, &order.IdProduct, &order.IdOrder, &order.Price, &order.Quantity, &order.TotalPrice, &order.OrderedDate, &order.ExpiredDate, &order.StatusPayment)
		helper.DoPanicIfError(err)
	}
}
