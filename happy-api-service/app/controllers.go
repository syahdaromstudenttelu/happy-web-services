package app

import (
	"happy-api-service/controller"
	"happy-api-service/util"

	"github.com/go-playground/validator/v10"
)

type AppControllers struct {
	ProductController  controller.ProductController
	UserController     controller.UserController
	FeedbackController controller.FeedbackController
	OrderController    controller.OrderController
	RegisterController controller.RegisterController
	LoginController    controller.LoginController
	LogoutController   controller.LogoutController
	AuthController     controller.AuthController
}

func GetAppControllers(config *util.Config, validate *validator.Validate) *AppControllers {
	return &AppControllers{
		ProductController: &controller.ProductControllerImpl{
			Config: config,
		},
		UserController: &controller.UserControllerImpl{
			Config: config,
		},
		FeedbackController: &controller.FeedbackControllerImpl{
			Config: config,
		},
		OrderController: &controller.OrderControllerImpl{
			Config: config,
		},
		RegisterController: &controller.RegisterControllerImpl{
			Config: config,
		},
		LoginController: &controller.LoginControllerImpl{
			Config:   config,
			Validate: validate,
		},
		LogoutController: &controller.LogoutControllerImpl{
			Config: config,
		},
		AuthController: &controller.AuthControllerImpl{
			Config: config,
		},
	}
}
