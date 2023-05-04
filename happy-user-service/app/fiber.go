package app

import (
	"happy-user-service/controller"
	"happy-user-service/exception"
	"happy-user-service/model/web"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewFiber(userController controller.UserController) *fiber.App {
	fiberApp := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if _, isOk := err.(validator.ValidationErrors); isOk {
				webResponse := web.WebResponse[string]{
					Code:   fiber.StatusNotAcceptable,
					Status: "failed",
					Data:   "request body is not valid",
				}
				return ctx.Status(fiber.StatusNotAcceptable).JSON(webResponse)
			}

			if duplicateAccountError, isOk := err.(*exception.DuplicateAccountError); isOk {
				var errMsg string

				if duplicateAccountError.Error() == "username_duplicate" {
					errMsg = "username has been used"
				}
				if duplicateAccountError.Error() == "email_duplicate" {
					errMsg = "email has been used"
				}

				webResponse := web.WebResponse[string]{
					Code:   fiber.StatusConflict,
					Status: "failed",
					Data:   errMsg,
				}
				return ctx.Status(fiber.StatusConflict).JSON(webResponse)
			}

			if userNotFoundError, isOk := err.(*exception.UserNotFoundError); isOk {
				webResponse := web.WebResponse[string]{
					Code:   fiber.StatusNotFound,
					Status: "failed",
					Data:   userNotFoundError.Error(),
				}

				return ctx.Status(fiber.StatusNotFound).JSON(webResponse)
			}

			webResponse := web.WebResponse[any]{
				Code:   fiber.StatusInternalServerError,
				Status: "failed",
				Data:   err.Error(),
			}

			return ctx.Status(fiber.StatusInternalServerError).JSON(webResponse)
		},
	})

	fiberApp.Use(recover.New())

	fiberApp.Get("/users", userController.FindAll)
	fiberApp.Get("/users/userId/:userId", userController.FindById)
	fiberApp.Get("/users/:username", userController.FindByUserName)
	fiberApp.Post("/users", userController.Create)

	return fiberApp
}
