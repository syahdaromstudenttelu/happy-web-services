package main

import (
	"happy-user-service/app"
	"happy-user-service/controller"
	"happy-user-service/helper"
	"happy-user-service/repository"
	"happy-user-service/service"
	"happy-user-service/util"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := util.LoadConfig(".")
	helper.DoPanicIfError(err)

	db := app.NewDB(config.DBDriver, config.DBSource)
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	fiberApp := app.NewFiber(userController)
	err = fiberApp.Listen(config.ServerAddr)
	helper.DoPanicIfError(err)
}
