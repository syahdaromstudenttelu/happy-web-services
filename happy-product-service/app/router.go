package app

import (
	"happy-product-service/controller"
	"happy-product-service/helper"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(productController controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/products", productController.FindAll)
	router.GET("/products/:productId", productController.FindById)
	router.PUT("/products", productController.Update)

	router.PanicHandler = helper.HttpRouterPanicHandler

	return router
}
