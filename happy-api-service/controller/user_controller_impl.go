package controller

import (
	"encoding/json"
	"fmt"
	"happy-api-service/helper"
	"happy-api-service/model/web"
	"happy-api-service/util"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct {
	Config *util.Config
}

func NewUserController(config *util.Config) UserController {
	return &UserControllerImpl{
		Config: config,
	}
}

func (controller *UserControllerImpl) FindByUserName(fiberCtx *fiber.Ctx) error {
	username := fiberCtx.Params("username")
	httpUrl := fmt.Sprintf("%s/users/%s", controller.Config.HappyUserServiceUrl, username)
	response, err := http.Get(httpUrl)
	helper.DoPanicIfError(err)

	webResponse := web.WebResponse[any]{}
	err = json.NewDecoder(response.Body).Decode(&webResponse)
	helper.DoPanicIfError(err)

	return fiberCtx.JSON(webResponse)
}

func (controller *UserControllerImpl) FindAll(fiberCtx *fiber.Ctx) error {
	httpUrl := fmt.Sprintf("%s/users", controller.Config.HappyUserServiceUrl)
	response, err := http.Get(httpUrl)
	helper.DoPanicIfError(err)

	webResponse := web.WebResponse[[]web.UserWebResponse]{}
	err = json.NewDecoder(response.Body).Decode(&webResponse)
	helper.DoPanicIfError(err)

	return fiberCtx.JSON(webResponse)
}
