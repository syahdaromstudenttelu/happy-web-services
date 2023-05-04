package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"happy-api-service/exception"
	"happy-api-service/helper"
	"happy-api-service/model/web"
	"happy-api-service/util"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type FeedbackControllerImpl struct {
	Config *util.Config
}

func NewFeedbackController(config *util.Config) FeedbackController {
	return &FeedbackControllerImpl{
		Config: config,
	}
}

func (controller *FeedbackControllerImpl) Create(fiberCtx *fiber.Ctx) error {
	reqBody := fiberCtx.Request().Body()
	idUserStr := string(fiberCtx.Request().Header.Cookie("id_user"))
	idUserInt, err := strconv.Atoi(idUserStr)
	helper.DoPanicIfError(err)

	feedbackUserRequest := web.FeedbackUserRequest{}
	err = json.Unmarshal(reqBody, &feedbackUserRequest)
	helper.DoPanicIfError(err)

	httpUrlPutOrderFeedbackDone := fmt.Sprintf("%s/orders/feedbackDone/%s", controller.Config.HappyOrderServiceUrl, feedbackUserRequest.IdOrder)
	helper.CreateHttpRequestService(http.MethodPut, httpUrlPutOrderFeedbackDone, nil)

	feedbackCreateRequest := web.FeedbackCreateRequest{
		IdUser:    uint(idUserInt),
		IdProduct: feedbackUserRequest.IdProduct,
		Feedback:  feedbackUserRequest.Feedback,
	}

	feedbackCreateReqMarshalled, err := json.Marshal(feedbackCreateRequest)
	helper.DoPanicIfError(err)

	httpUrlPostFeedback := fmt.Sprintf("%s/feedbacks", controller.Config.HappyFeedbackServiceUrl)
	feedbackSvcRes := helper.CreateHttpRequestService(http.MethodPost, httpUrlPostFeedback, bytes.NewBuffer(feedbackCreateReqMarshalled))

	return fiberCtx.JSON(feedbackSvcRes)
}

func (controller *FeedbackControllerImpl) FindByProductId(fiberCtx *fiber.Ctx) error {
	productId := fiberCtx.Params("productId")
	productIdInt, err := strconv.Atoi(productId)

	if err != nil {
		panic(exception.NewUrlParamError("url param is not valid"))
	}

	httpUrlGet := fmt.Sprintf("%s/feedbacks/%d", controller.Config.HappyFeedbackServiceUrl, uint(productIdInt))
	feedbacksSvcRes := helper.CreateHttpRequestService(http.MethodGet, httpUrlGet, nil)

	return fiberCtx.JSON(feedbacksSvcRes)
}
