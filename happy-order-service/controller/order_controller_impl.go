package controller

import (
	"happy-order-service/helper"
	"happy-order-service/model/web"
	"happy-order-service/service"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

type OrderControllerImpl struct {
	OrderService service.OrderService
}

func NewOrderController(orderService service.OrderService) OrderController {
	return &OrderControllerImpl{
		OrderService: orderService,
	}
}

func (controller *OrderControllerImpl) Create(fiberCtx *fiber.Ctx) error {
	orderCreateRequest := web.OrderCreateRequest{}
	helper.ReadFromRequestBody(fiberCtx.Request(), &orderCreateRequest)

	orderResponse := controller.OrderService.Create(fiberCtx.Context(), orderCreateRequest)

	webResponse := web.WebResponse[fiber.Map]{
		Code:   fiber.StatusCreated,
		Status: "success",
		Data: fiber.Map{
			"orderId": orderResponse.IdOrder,
		},
	}

	return helper.WriteToResponseBody(fiberCtx, webResponse, fiber.StatusCreated)
}

func (controller *OrderControllerImpl) UpdateStatusPayment(fiberCtx *fiber.Ctx) error {
	orderId := fiberCtx.Params("orderId")

	controller.OrderService.UpdateStatusPayment(fiberCtx.Context(), orderId)

	webResponse := web.WebResponse[fiber.Map]{
		Code:   fiber.StatusOK,
		Status: "success",
		Data: fiber.Map{
			"statusPayment": true,
		},
	}

	return helper.WriteToResponseBody(fiberCtx, webResponse, fiber.StatusOK)
}

func (controller *OrderControllerImpl) UpdateFeedbackDone(fiberCtx *fiber.Ctx) error {
	orderId := fiberCtx.Params("orderId")

	controller.OrderService.UpdateFeedbackDone(fiberCtx.Context(), orderId)

	webResponse := web.WebResponse[fiber.Map]{
		Code:   fiber.StatusOK,
		Status: "success",
		Data: fiber.Map{
			"feedbackDone": true,
		},
	}

	return helper.WriteToResponseBody(fiberCtx, webResponse, fiber.StatusOK)
}

func (controller *OrderControllerImpl) Delete(fiberCtx *fiber.Ctx) error {
	userId := fiberCtx.Params("userId")
	userIdInt, err := strconv.Atoi(userId)
	helper.DoPanicIfError(err)
	orderId := fiberCtx.Params("orderId")

	controller.OrderService.Delete(fiberCtx.Context(), uint(userIdInt), orderId)

	webResponse := web.WebResponse[fiber.Map]{
		Code:   fiber.StatusOK,
		Status: "success",
		Data:   nil,
	}

	return helper.WriteToResponseBody(fiberCtx, webResponse, fiber.StatusOK)
}

func (controller *OrderControllerImpl) FindByIdOrder(fiberCtx *fiber.Ctx) error {
	orderId := fiberCtx.Params("orderId")

	orderResponse := controller.OrderService.FindByIdOrder(fiberCtx.Context(), orderId)

	webResponse := web.WebResponse[web.OrderResponse]{
		Code:   fiber.StatusOK,
		Status: "success",
		Data:   orderResponse,
	}

	return helper.WriteToResponseBody(fiberCtx, webResponse, fiber.StatusOK)
}

func (controller *OrderControllerImpl) FindByIdUserAndIdOrder(fiberCtx *fiber.Ctx) error {
	userId := fiberCtx.Params("userId")
	userIdInt, err := strconv.Atoi(userId)
	helper.DoPanicIfError(err)
	orderId := fiberCtx.Params("orderId")

	orderResponse := controller.OrderService.FindByIdUserAndIdOrder(fiberCtx.Context(), uint(userIdInt), orderId)

	webResponse := web.WebResponse[web.OrderResponse]{
		Code:   fiber.StatusOK,
		Status: "success",
		Data:   orderResponse,
	}

	return helper.WriteToResponseBody(fiberCtx, webResponse, fiber.StatusOK)
}

func (controller *OrderControllerImpl) FindAllByIdUser(fiberCtx *fiber.Ctx) error {
	userId := fiberCtx.Params("userId")
	userIdInt, err := strconv.Atoi(userId)
	helper.DoPanicIfError(err)

	ordersResponse := controller.OrderService.FindAllByIdUser(fiberCtx.Context(), uint(userIdInt))

	webResponse := web.WebResponse[[]web.OrderResponse]{
		Code:   fiber.StatusOK,
		Status: "success",
		Data:   ordersResponse,
	}

	return helper.WriteToResponseBody(fiberCtx, webResponse, fiber.StatusOK)
}

func (controller *OrderControllerImpl) FindAll(fiberCtx *fiber.Ctx) error {
	ordersResponse := controller.OrderService.FindAll(fiberCtx.Context())

	webResponse := web.WebResponse[[]web.OrderResponse]{
		Code:   fiber.StatusOK,
		Status: "success",
		Data:   ordersResponse,
	}

	return helper.WriteToResponseBody(fiberCtx, webResponse, fiber.StatusOK)
}
