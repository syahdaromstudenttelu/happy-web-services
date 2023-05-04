package controller

import (
	"happy-feedback-service/helper"
	"happy-feedback-service/model/web"
	"happy-feedback-service/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type FeedbackControllerImpl struct {
	FeedbackService service.FeedbackService
}

func NewFeedbackController(feedbackService service.FeedbackService) FeedbackController {
	return &FeedbackControllerImpl{
		FeedbackService: feedbackService,
	}
}

func (controller *FeedbackControllerImpl) Create(fiberCtx *fiber.Ctx) error {
	feedbackCreateRequest := web.FeedbackCreateRequest{}
	helper.ReadFromRequestBody(fiberCtx.Request(), &feedbackCreateRequest)

	feedbackResponse := controller.FeedbackService.Create(fiberCtx.Context(), feedbackCreateRequest)

	webResponse := web.WebResponse[any]{
		Code:   fiber.StatusCreated,
		Status: "status",
		Data: fiber.Map{
			"feedbackId": feedbackResponse.Id,
		},
	}

	return helper.WriteToResponseBody(fiberCtx, &webResponse, fiber.StatusCreated)
}

func (controller *FeedbackControllerImpl) FindByIdProduct(fiberCtx *fiber.Ctx) error {
	productIdStr := fiberCtx.Params("productId")
	productId, err := strconv.Atoi(productIdStr)
	helper.DoPanicIfError(err)

	feedbacksResponse := controller.FeedbackService.FindByIdProduct(fiberCtx.Context(), uint(productId))

	webResponse := web.WebResponse[[]web.FeedbackResponse]{
		Code:   fiber.StatusOK,
		Status: "success",
		Data:   feedbacksResponse,
	}

	return helper.WriteToResponseBody(fiberCtx, &webResponse, fiber.StatusOK)
}

func (controller *FeedbackControllerImpl) FindAll(fiberCtx *fiber.Ctx) error {
	feedbacksResponse := controller.FeedbackService.FindAll(fiberCtx.Context())

	webResponse := web.WebResponse[[]web.FeedbackResponse]{
		Code:   fiber.StatusOK,
		Status: "success",
		Data:   feedbacksResponse,
	}

	return helper.WriteToResponseBody(fiberCtx, &webResponse, fiber.StatusOK)
}
