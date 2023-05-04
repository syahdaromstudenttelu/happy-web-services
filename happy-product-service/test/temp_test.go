package test

import (
	"context"
	"fmt"
	"happy-product-service/app"
	"happy-product-service/helper"
	"happy-product-service/util"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestTemp(t *testing.T) {
	defer func() {
		err := recover()
		fmt.Println(strings.Contains(err.(error).Error(), "3819"))
		fmt.Println(err.(error).Error())
	}()

	config, err := util.LoadConfig("../")
	helper.DoPanicIfError(err)

	db := app.NewDB(config.DBDriver, config.DBSource)
	tx, err := db.Begin()
	helper.DoPanicIfError(err)

	ctx := context.Background()

	sqlExec := "UPDATE product set reservation = 500 WHERE id = 1"
	_, err = tx.ExecContext(ctx, sqlExec)
	helper.DoPanicIfError(err)
}
