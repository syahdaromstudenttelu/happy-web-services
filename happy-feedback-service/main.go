package main

import (
	"happy-feedback-service/app"
	"happy-feedback-service/controller"
	"happy-feedback-service/helper"
	"happy-feedback-service/repository"
	"happy-feedback-service/service"
	"happy-feedback-service/util"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := util.LoadConfig(".")
	helper.DoPanicIfError(err)

	db := app.NewDB(config.DBDriver, config.DBSource)
	validate := validator.New()
	feedbackRepository := repository.NewFeedbackRepository()
	feedbackService := service.NewFeedbackService(feedbackRepository, db, validate)
	feedbackController := controller.NewFeedbackController(feedbackService)

	fiberApp := app.NewFiber(feedbackController)
	err = fiberApp.Listen(config.ServerAddr)
	helper.DoPanicIfError(err)
}
