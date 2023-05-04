package main

import (
	"happy-product-service/app"
	"happy-product-service/controller"
	"happy-product-service/helper"
	"happy-product-service/repository"
	"happy-product-service/service"
	"happy-product-service/util"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := util.LoadConfig(".")
	helper.DoPanicIfError(err)

	db := app.NewDB(config.DBDriver, config.DBSource)
	validate := validator.New()
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)
	router := app.NewRouter(productController)

	server := http.Server{
		Addr:    config.ServerAddr,
		Handler: router,
	}

	err = server.ListenAndServe()
	helper.DoPanicIfError(err)
}
