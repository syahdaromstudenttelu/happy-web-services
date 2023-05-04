package app

import (
	"happy-order-service/controller"
	"happy-order-service/exception"
	"happy-order-service/model/web"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewFiber(orderController controller.OrderController) *fiber.App {
	fiberApp := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if reqMalformedError, isOk := err.(*exception.ReqBodyMalformedError); isOk {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"code":    fiber.StatusBadRequest,
					"status":  "failed",
					"message": reqMalformedError.Error(),
				})
			}

			if orderNotFoundError, isOk := err.(*exception.OrderNotFoundError); isOk {
				webResponse := web.WebResponse[string]{
					Code:   fiber.StatusNotFound,
					Status: "failed",
					Data:   orderNotFoundError.Error(),
				}

				return ctx.Status(fiber.StatusNotFound).JSON(webResponse)
			}

			webResponse := web.WebResponse[any]{
				Code:   fiber.StatusInternalServerError,
				Status: "failed",
				Data:   nil,
			}

			return ctx.Status(fiber.StatusInternalServerError).JSON(webResponse)
		},
	})

	fiberApp.Use(recover.New())

	fiberApp.Get("/orders", orderController.FindAll)
	fiberApp.Get("/orders/orderId/:orderId", orderController.FindByIdOrder)
	fiberApp.Get("/orders/:userId", orderController.FindAllByIdUser)
	fiberApp.Get("/orders/userId::userId/orderId::orderId", orderController.FindByIdUserAndIdOrder)
	fiberApp.Post("/orders", orderController.Create)
	fiberApp.Put("/orders/statusPayment/:orderId", orderController.UpdateStatusPayment)
	fiberApp.Put("/orders/feedbackDone/:orderId", orderController.UpdateFeedbackDone)
	fiberApp.Delete("/orders/userId::userId/orderId::orderId", orderController.Delete)

	return fiberApp
}
