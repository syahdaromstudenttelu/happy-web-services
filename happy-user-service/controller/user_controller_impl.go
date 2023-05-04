package controller

import (
	"happy-user-service/helper"
	"happy-user-service/model/web"
	"happy-user-service/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Create(fiberCtx *fiber.Ctx) error {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(fiberCtx.Request(), &userCreateRequest)

	userResponse := controller.UserService.Create(fiberCtx.Context(), userCreateRequest)

	webResponse := web.WebResponse[web.UserResponse]{
		Code:   201,
		Status: "CREATED",
		Data:   userResponse,
	}

	return helper.WriteToResponseBody(fiberCtx, fiber.StatusCreated, &webResponse)
}

func (controller *UserControllerImpl) FindById(fiberCtx *fiber.Ctx) error {
	userId := fiberCtx.Params("userId")
	userIdInt, err := strconv.Atoi(userId)
	helper.DoPanicIfError(err)

	userResponse := controller.UserService.FindById(fiberCtx.Context(), uint(userIdInt))

	webResponse := web.WebResponse[web.UserResponse]{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return helper.WriteToResponseBody(fiberCtx, fiber.StatusOK, &webResponse)
}

func (controller *UserControllerImpl) FindByUserName(fiberCtx *fiber.Ctx) error {
	username := fiberCtx.Params("username")

	userResponse := controller.UserService.FindByUserName(fiberCtx.Context(), username)

	webResponse := web.WebResponse[web.UserResponse]{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	return helper.WriteToResponseBody(fiberCtx, fiber.StatusOK, &webResponse)
}

func (controller *UserControllerImpl) FindAll(fiberCtx *fiber.Ctx) error {
	usersResponse := controller.UserService.FindAll(fiberCtx.Context())

	webResponse := web.WebResponse[[]web.UserResponse]{
		Code:   200,
		Status: "OK",
		Data:   usersResponse,
	}

	return helper.WriteToResponseBody(fiberCtx, fiber.StatusOK, &webResponse)
}
