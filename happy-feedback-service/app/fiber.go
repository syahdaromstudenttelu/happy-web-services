package app

import (
	"happy-feedback-service/controller"
	"happy-feedback-service/exception"
	"happy-feedback-service/model/web"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewFiber(feedbackController controller.FeedbackController) *fiber.App {
	fiberApp := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if reqMalformedError, isOk := err.(*exception.ReqBodyMalformedError); isOk {
				webResponse := web.WebResponse[string]{
					Code:   fiber.StatusBadRequest,
					Status: "failed",
					Data:   reqMalformedError.Error(),
				}
				return ctx.Status(fiber.StatusBadRequest).JSON(webResponse)
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

	fiberApp.Get("/feedbacks", feedbackController.FindAll)
	fiberApp.Get("/feedbacks/:productId", feedbackController.FindByIdProduct)
	fiberApp.Post("/feedbacks", feedbackController.Create)

	return fiberApp
}
