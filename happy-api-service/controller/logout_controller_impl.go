package controller

import (
	"happy-api-service/model/web"
	"happy-api-service/util"

	"github.com/gofiber/fiber/v2"
)

type LogoutControllerImpl struct {
	Config *util.Config
}

func NewLogoutController(config *util.Config) LogoutController {
	return &LogoutControllerImpl{
		Config: config,
	}
}

func (controller *LogoutControllerImpl) Logout(fiberCtx *fiber.Ctx) error {
	jwtRmCookie := &fiber.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
	}

	userIdRmCookie := &fiber.Cookie{
		Name:     "id_user",
		Path:     "/",
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
	}

	fiberCtx.Cookie(jwtRmCookie)
	fiberCtx.Cookie(userIdRmCookie)

	webResponse := web.WebResponse[string]{
		Code:   fiber.StatusOK,
		Status: "success",
		Data:   "logout is successfully",
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(webResponse)
}
