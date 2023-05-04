package controller

import (
	"bytes"
	"fmt"
	"happy-api-service/helper"
	"happy-api-service/model/web"
	"happy-api-service/util"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type RegisterControllerImpl struct {
	Config *util.Config
}

func NewRegisterController(config *util.Config) RegisterController {
	return &RegisterControllerImpl{
		Config: config,
	}
}

func (controller *RegisterControllerImpl) Register(fiberCtx *fiber.Ctx) error {
	reqBody := fiberCtx.Request().Body()

	httpUrlPost := fmt.Sprintf("%s/users", controller.Config.HappyUserServiceUrl)
	userSvcRes := helper.CreateHttpRequestService(http.MethodPost, httpUrlPost, bytes.NewBuffer(reqBody))

	if userSvcRes.Status == "failed" {
		return fiberCtx.Status(fiber.StatusConflict).JSON(userSvcRes)
	}

	webResponse := web.WebResponse[string]{
		Code:   201,
		Status: "success",
		Data:   "account has successfully created",
	}

	return fiberCtx.Status(fiber.StatusCreated).JSON(webResponse)
}
