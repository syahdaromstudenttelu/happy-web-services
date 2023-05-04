package app

import (
	"happy-api-service/exception"
	"happy-api-service/middleware"
	"happy-api-service/model/web"
	"happy-api-service/util"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewFiber(config *util.Config, validate *validator.Validate) *fiber.App {
	appControllers := GetAppControllers(config, validate)

	fiberApp := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if urlParamError, isOk := err.(*exception.UrlParamError); isOk {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"code":    fiber.StatusBadRequest,
					"status":  "failed",
					"message": urlParamError.Error(),
				})
			}

			webResponse := web.WebResponse[any]{
				Code:   fiber.StatusInternalServerError,
				Status: "failed",
				Data:   nil,
			}

			return ctx.Status(fiber.StatusInternalServerError).JSON(webResponse)
		},
	})

	allowOrigins := strings.Join(strings.Split(config.AllowOrigins, ","), ", ")

	fiberApp.Use(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowCredentials: true,
	}))

	fiberApp.Use(recover.New())

	fiberApi := fiberApp.Group("/api")

	// # Non Auth Endpoints
	fiberApi.Post("/login", appControllers.LoginController.Login)
	fiberApi.Post("/register", appControllers.RegisterController.Register)
	fiberApi.Get("/logout", appControllers.LogoutController.Logout)

	// Products API
	fiberApi.Get("/products", appControllers.ProductController.FindAll)
	fiberApi.Get("/products/:productId", appControllers.ProductController.FindById)

	// # Auth Endpoints
	fiberApi.Use(middleware.JwtMiddleware)

	// Auth Authentication API
	fiberApi.Get("/auth", appControllers.AuthController.Auth)

	// Auth Feedbacks API
	fiberApi.Post("/feedbacks", appControllers.FeedbackController.Create)

	// Auth Orders API
	fiberApi.Get("/orders/orderByIdUser", appControllers.OrderController.FindByUserId)
	fiberApi.Get("/orders/:orderId", appControllers.OrderController.FindByUserAndOrderId)
	fiberApi.Post("/orders", appControllers.OrderController.Create)
	fiberApi.Delete("/orders/:orderId", appControllers.OrderController.Delete)

	return fiberApp
}
