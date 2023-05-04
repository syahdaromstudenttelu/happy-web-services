package controller

import (
	"fmt"
	"happy-api-service/helper"
	"happy-api-service/model/web"
	"happy-api-service/util"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AuthControllerImpl struct {
	Config *util.Config
}

func NewAuthController(config *util.Config) AuthController {
	return &AuthControllerImpl{
		Config: config,
	}
}

func (controller *AuthControllerImpl) Auth(fiberCtx *fiber.Ctx) error {
	idUser := string(fiberCtx.Request().Header.Cookie("id_user"))
	idUserInt, err := strconv.Atoi(idUser)
	helper.DoPanicIfError(err)

	httpUrlGetUser := fmt.Sprintf("%s/users/userId/%d", controller.Config.HappyUserServiceUrl, idUserInt)
	userSvcRes := helper.CreateHttpRequestService(http.MethodGet, httpUrlGetUser, nil)

	if userSvcRes.Status == "failed" {
		return fiberCtx.JSON(userSvcRes)
	}

	userSvcResData := web.UserServiceResponse{}
	helper.GetServiceDataResponse(&userSvcRes.Data, &userSvcResData)

	userWebResponse := web.UserWebResponse{}
	helper.JoinResponse(&userSvcResData, &userWebResponse)

	webResponse := web.WebResponse[web.UserWebResponse]{
		Code:   fiber.StatusOK,
		Status: "success",
		Data:   userWebResponse,
	}

	fiberCtx.Response().Header.SetStatusCode(fiber.StatusOK)
	return fiberCtx.JSON(webResponse)
}
