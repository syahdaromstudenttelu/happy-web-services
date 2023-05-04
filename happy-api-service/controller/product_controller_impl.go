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
	"time"

	"github.com/gofiber/fiber/v2"
)

type ProductControllerImpl struct {
	Config *util.Config
}

func NewProductController(config *util.Config) ProductController {
	return &ProductControllerImpl{
		Config: config,
	}
}

func getAllUsersAndFeedbacksSvc(controller *ProductControllerImpl) (usersSvcResData []web.UserServiceResponse, feedbacksSvcResData []web.FeedbackServiceResponse) {
	// Get All Users
	httpUrlGetUsers := fmt.Sprintf("%s/users", controller.Config.HappyUserServiceUrl)

	svcUsersRes := helper.CreateHttpRequestService(http.MethodGet, httpUrlGetUsers, nil)

	usersSvcResData = []web.UserServiceResponse{}
	helper.GetServiceDataResponse(&svcUsersRes.Data, &usersSvcResData)

	// Get All Feedbacks
	httpUrlGetFeedback := fmt.Sprintf("%s/feedbacks", controller.Config.HappyFeedbackServiceUrl)

	svcFeedbacksRes := helper.CreateHttpRequestService(http.MethodGet, httpUrlGetFeedback, nil)

	feedbacksSvcResData = []web.FeedbackServiceResponse{}
	helper.GetServiceDataResponse(&svcFeedbacksRes.Data, &feedbacksSvcResData)

	return usersSvcResData, feedbacksSvcResData
}

func (controller *ProductControllerImpl) Update(fiberCtx *fiber.Ctx) error {
	reqBody := fiberCtx.Request().Body()
	httpUrl := fmt.Sprintf("%s/products", controller.Config.HappyProductServiceUrl)

	svcProductUpdateRes := helper.CreateHttpRequestService(http.MethodPut, httpUrl, bytes.NewBuffer(reqBody))

	return fiberCtx.JSON(svcProductUpdateRes)
}

func (controller *ProductControllerImpl) FindById(fiberCtx *fiber.Ctx) error {
	productIdStr := fiberCtx.Params("productId")
	productIdInt, err := strconv.Atoi(productIdStr)

	if err != nil {
		panic(exception.NewUrlParamError("url param invalid"))
	}

	productWebResponse := web.ProductWebResponse{}

	httpUrlGetProduct := fmt.Sprintf("%s/products/%d", controller.Config.HappyProductServiceUrl, productIdInt)

	serviceProductResponse := helper.CreateHttpRequestService(http.MethodGet, httpUrlGetProduct, nil)

	if serviceProductResponse.Status == "failed" {
		return fiberCtx.Status(fiber.StatusNotFound).JSON(serviceProductResponse)
	}

	productSvcResData := web.ProductServiceResponse{}
	helper.GetServiceDataResponse(&serviceProductResponse.Data, &productSvcResData)
	helper.JoinResponse(&productSvcResData, &productWebResponse)

	usersSvcResData, feedbacksSvcResData := getAllUsersAndFeedbacksSvc(controller)

	for _, feedbackSvcResData := range feedbacksSvcResData {
		if feedbackSvcResData.IdProduct == productWebResponse.Id {
			feedback := web.FeedbackWebResponse{}
			helper.JoinResponse(&feedbackSvcResData, &feedback)

			for _, userSvcResData := range usersSvcResData {
				if feedbackSvcResData.IdUser == userSvcResData.Id {
					feedback.FullName = userSvcResData.FullName
				}
			}

			productWebResponse.Feedbacks = append(productWebResponse.Feedbacks, feedback)
		}
	}

	webResponse := web.WebResponse[web.ProductWebResponse]{
		Code:   fiber.StatusOK,
		Status: "success",
		Data:   productWebResponse,
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *ProductControllerImpl) FindAll(fiberCtx *fiber.Ctx) error {
	// Get All Orders
	httpUrlGetOrders := fmt.Sprintf("%s/orders", controller.Config.HappyOrderServiceUrl)

	svcOrdersRes := helper.CreateHttpRequestService(http.MethodGet, httpUrlGetOrders, nil)

	svcOrdersResData := []web.OrderServiceResponse{}

	helper.GetServiceDataResponse(&svcOrdersRes.Data, &svcOrdersResData)

	for _, order := range svcOrdersResData {
		currentTime := time.Now()
		expTime, err := time.Parse(time.RFC3339, order.ExpiredDate)
		helper.DoPanicIfError(err)

		orderIsExpired := expTime.Before(currentTime)

		if orderIsExpired {
			productUpdateRequest := web.ProductUpdateRequest{
				Id:              order.IdProduct,
				IsOrderReject:   true,
				ProductReserved: order.Quantity,
			}

			productUpdateReqBytes, err := json.Marshal(productUpdateRequest)
			helper.DoPanicIfError(err)

			httpUrlPutProduct := fmt.Sprintf("%s/products", controller.Config.HappyProductServiceUrl)
			helper.CreateHttpRequestService(http.MethodPut, httpUrlPutProduct, bytes.NewBuffer(productUpdateReqBytes))
		}
	}

	// Get All Products
	productsWebResponse := []web.ProductWebResponse{}

	httpUrlGetProducts := fmt.Sprintf("%s/products", controller.Config.HappyProductServiceUrl)

	serviceProductsResponse := helper.CreateHttpRequestService(http.MethodGet, httpUrlGetProducts, nil)

	if serviceProductsResponse.Status == "failed" {
		return fiberCtx.JSON(serviceProductsResponse)
	}

	productsSvcResData := []web.ProductServiceResponse{}
	helper.GetServiceDataResponse(&serviceProductsResponse.Data, &productsSvcResData)

	for _, product := range productsSvcResData {
		productWebResponse := web.ProductWebResponse{}
		helper.JoinResponse(&product, &productWebResponse)

		productsWebResponse = append(productsWebResponse, productWebResponse)
	}

	usersSvcResData, feedbacksSvcResData := getAllUsersAndFeedbacksSvc(controller)

	for idx, product := range productsWebResponse {
		for _, feedbackSvcResData := range feedbacksSvcResData {
			if feedbackSvcResData.IdProduct == product.Id {
				feedback := web.FeedbackWebResponse{}
				helper.JoinResponse(&feedbackSvcResData, &feedback)

				for _, userSvcResData := range usersSvcResData {
					if feedbackSvcResData.IdUser == userSvcResData.Id {
						feedback.FullName = userSvcResData.FullName
					}
				}

				productsWebResponse[idx].Feedbacks = append(productsWebResponse[idx].Feedbacks, feedback)
			}
		}
	}

	webResponse := web.WebResponse[[]web.ProductWebResponse]{
		Code:   fiber.StatusOK,
		Status: "success",
		Data:   productsWebResponse,
	}

	return fiberCtx.JSON(webResponse)
}
