package main

import (
	"happy-api-service/app"
	"happy-api-service/helper"
	"happy-api-service/util"

	"github.com/go-playground/validator/v10"
)

func main() {
	config, err := util.LoadConfig(".")
	helper.DoPanicIfError(err)

	validate := validator.New()
	fiberApp := app.NewFiber(&config, validate)

	err = fiberApp.Listen(config.ServerAddr)
	helper.DoPanicIfError(err)
}
