package main

import (
	"happy-order-service/app"
	"happy-order-service/controller"
	"happy-order-service/helper"
	"happy-order-service/repository"
	"happy-order-service/service"
	"happy-order-service/util"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := util.LoadConfig(".")
	helper.DoPanicIfError(err)

	db := app.NewDB(config.DBDriver, config.DBSource)
	validate := validator.New()
	orderRepository := repository.NewOrderRepository()
	orderService := service.NewOrderService(orderRepository, db, validate)
	orderController := controller.NewOrderController(orderService)

	fiberApp := app.NewFiber(orderController)
	err = fiberApp.Listen(config.ServerAddr)
	helper.DoPanicIfError(err)
}
