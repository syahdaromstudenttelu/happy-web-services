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

type OrderControllerImpl struct {
	Config *util.Config
}

func NewOrderController(config *util.Config) OrderController {
	return &OrderControllerImpl{
		Config: config,
	}
}

func (controller *OrderControllerImpl) Create(fiberCtx *fiber.Ctx) error {
	reqBody := fiberCtx.Request().Body()
	idUserStr := string(fiberCtx.Request().Header.Cookie("id_user"))
	idUserInt, err := strconv.Atoi(idUserStr)
	helper.DoPanicIfError(err)

	orderCreateRequest := web.OrderCreateRequest{}
	err = json.Unmarshal(reqBody, &orderCreateRequest)
	helper.DoPanicIfError(err)

	productUpdateRequest := web.ProductUpdateRequest{
		Id:              orderCreateRequest.IdProduct,
		IsOrder:         true,
		ProductReserved: orderCreateRequest.Quantity,
	}

	productUpdateMarshalled, err := json.Marshal(productUpdateRequest)
	helper.DoPanicIfError(err)

	httpUrlUpdateProduct := fmt.Sprintf("%s/products", controller.Config.HappyProductServiceUrl)
	productSvcUpdRes := helper.CreateHttpRequestService(http.MethodPut, httpUrlUpdateProduct, bytes.NewBuffer(productUpdateMarshalled))

	if productSvcUpdRes.Status == "failed" {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(productSvcUpdRes)
	}

	httpUrlGetProduct := fmt.Sprintf("%s/products/%d", controller.Config.HappyProductServiceUrl, orderCreateRequest.IdProduct)
	productSvcRes := helper.CreateHttpRequestService(http.MethodGet, httpUrlGetProduct, nil)

	if productSvcRes.Status == "failed" {
		return fiberCtx.Status(fiber.StatusNotFound).JSON(productSvcRes)
	}

	productSvcResData := web.ProductServiceResponse{}
	helper.JoinResponse(&productSvcRes.Data, &productSvcResData)

	orderCreateRequest.IdUser = uint(idUserInt)
	orderCreateRequest.Price = productSvcResData.ProductPrice

	marshalledOrderReq, err := json.Marshal(orderCreateRequest)
	helper.DoPanicIfError(err)

	httpUrlPostOrder := fmt.Sprintf("%s/orders", controller.Config.HappyOrderServiceUrl)
	orderSvcRes := helper.CreateHttpRequestService(http.MethodPost, httpUrlPostOrder, bytes.NewBuffer(marshalledOrderReq))

	return fiberCtx.Status(fiber.StatusCreated).JSON(orderSvcRes)
}

func (controller *OrderControllerImpl) UpdateByStatusPayment(fiberCtx *fiber.Ctx) error {
	orderId := fiberCtx.Params("orderId")

	httpUrlGetOrder := fmt.Sprintf("%s/orders/orderId/%s", controller.Config.HappyOrderServiceUrl, orderId)
	orderSvcGetRes := helper.CreateHttpRequestService(http.MethodGet, httpUrlGetOrder, nil)

	if orderSvcGetRes.Status == "failed" {
		fiberCtx.Response().Header.SetStatusCode(fiber.StatusNotFound)
		return fiberCtx.JSON(orderSvcGetRes)
	}

	orderSvcGetResData := web.OrderServiceResponse{}
	helper.GetServiceDataResponse(&orderSvcGetRes.Data, &orderSvcGetResData)

	productUpdateRequest := web.ProductUpdateRequest{
		Id:              orderSvcGetResData.IdProduct,
		IsPaid:          true,
		ProductReserved: orderSvcGetResData.Quantity,
	}

	productUpdReqMarshalled, err := json.Marshal(productUpdateRequest)
	helper.DoPanicIfError(err)

	httpUrlPutProduct := fmt.Sprintf("%s/products", controller.Config.HappyProductServiceUrl)
	helper.CreateHttpRequestService(http.MethodPut, httpUrlPutProduct, bytes.NewBuffer(productUpdReqMarshalled))

	httpUrlPutOrder := fmt.Sprintf("%s/orders/statusPayment/%s", controller.Config.HappyOrderServiceUrl, orderId)
	orderSvcPutRes := helper.CreateHttpRequestService(http.MethodPut, httpUrlPutOrder, nil)

	return fiberCtx.Status(fiber.StatusOK).JSON(orderSvcPutRes)
}

func (controller *OrderControllerImpl) Delete(fiberCtx *fiber.Ctx) error {
	idUserStr := string(fiberCtx.Request().Header.Cookie("id_user"))
	idUserInt, err := strconv.Atoi(idUserStr)
	helper.DoPanicIfError(err)
	orderId := fiberCtx.Params("orderId")

	if err != nil {
		panic(exception.NewUrlParamError("url param is not valid"))
	}

	httpUrlGetOrder := fmt.Sprintf("%s/orders/userId:%d/orderId:%s", controller.Config.HappyOrderServiceUrl, uint(idUserInt), orderId)
	orderSvcGetOrder := helper.CreateHttpRequestService(http.MethodGet, httpUrlGetOrder, nil)

	orderSvcGetOrderData := web.OrderServiceResponse{}
	helper.JoinResponse(&orderSvcGetOrder.Data, &orderSvcGetOrderData)

	productUpdateRequest := web.ProductUpdateRequest{
		Id:              orderSvcGetOrderData.IdProduct,
		IsOrderReject:   true,
		ProductReserved: orderSvcGetOrderData.Quantity,
	}

	productUpdateMarshalled, err := json.Marshal(productUpdateRequest)
	helper.DoPanicIfError(err)

	httpUrlUpdateProduct := fmt.Sprintf("%s/products", controller.Config.HappyProductServiceUrl)
	helper.CreateHttpRequestService(http.MethodPut, httpUrlUpdateProduct, bytes.NewBuffer(productUpdateMarshalled))

	httpUrlDelOrder := fmt.Sprintf("%s/orders/userId:%d/orderId:%s", controller.Config.HappyOrderServiceUrl, uint(idUserInt), orderId)
	orderSvcDel := helper.CreateHttpRequestService(http.MethodDelete, httpUrlDelOrder, nil)

	return fiberCtx.Status(fiber.StatusOK).JSON(orderSvcDel)
}

func (controller *OrderControllerImpl) FindByUserId(fiberCtx *fiber.Ctx) error {
	userIdStr := string(fiberCtx.Request().Header.Cookie("id_user"))
	userIdInt, err := strconv.Atoi(userIdStr)
	helper.DoPanicIfError(err)

	ordersWebResponse := []web.OrderWebResponse{}

	httpUrlGetOrders := fmt.Sprintf("%s/orders/%d", controller.Config.HappyOrderServiceUrl, uint(userIdInt))

	ordersSvcRes := helper.CreateHttpRequestService(http.MethodGet, httpUrlGetOrders, nil)

	if ordersSvcRes.Status == "failed" {
		return fiberCtx.Status(fiber.StatusNotFound).JSON(ordersSvcRes)
	}

	ordersSvcResData := []web.OrderServiceResponse{}
	helper.GetServiceDataResponse(&ordersSvcRes.Data, &ordersSvcResData)

	for _, order := range ordersSvcResData {
		ordersWebResponse = append(ordersWebResponse, web.OrderWebResponse{
			IdOrder:       order.IdOrder,
			IdProduct:     order.IdProduct,
			Price:         order.Price,
			Quantity:      order.Quantity,
			TotalPrice:    order.TotalPrice,
			OrderedDate:   order.OrderedDate,
			ExpiredDate:   order.ExpiredDate,
			StatusPayment: order.StatusPayment,
			FeedbackDone:  order.FeedbackDone,
		})
	}

	httpUrlGetProducts := fmt.Sprintf("%s/products", controller.Config.HappyProductServiceUrl)
	svcProductsRes := helper.CreateHttpRequestService(http.MethodGet, httpUrlGetProducts, nil)

	if svcProductsRes.Status == "failed" {
		return fiberCtx.Status(fiber.StatusNotFound).JSON(svcProductsRes)
	}

	productsSvcResData := []web.ProductServiceResponse{}
	helper.GetServiceDataResponse(&svcProductsRes.Data, &productsSvcResData)

	for idx, order := range ordersWebResponse {
		for _, product := range productsSvcResData {
			if product.Id == order.IdProduct {
				ordersWebResponse[idx].Brand = product.Brand
				ordersWebResponse[idx].Type = product.Type
				ordersWebResponse[idx].Name = product.Name
				ordersWebResponse[idx].PriceName = product.PriceName

				break
			}
		}
	}

	webResponse := web.WebResponse[[]web.OrderWebResponse]{
		Code:   fiber.StatusOK,
		Status: "success",
		Data:   ordersWebResponse,
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *OrderControllerImpl) FindByUserAndOrderId(fiberCtx *fiber.Ctx) error {
	orderId := fiberCtx.Params("orderId")
	userIdStr := string(fiberCtx.Request().Header.Cookie("id_user"))
	userIdInt, err := strconv.Atoi(userIdStr)
	helper.DoPanicIfError(err)

	orderWebResponse := web.OrderWebResponse{}

	httpUrlGetOrders := fmt.Sprintf("%s/orders/userId:%d/orderId:%s", controller.Config.HappyOrderServiceUrl, uint(userIdInt), orderId)
	orderSvcRes := helper.CreateHttpRequestService(http.MethodGet, httpUrlGetOrders, nil)

	if orderSvcRes.Status == "failed" {
		return fiberCtx.Status(fiber.StatusNotFound).JSON(orderSvcRes)
	}

	orderSvcResData := web.OrderServiceResponse{}
	helper.GetServiceDataResponse(&orderSvcRes.Data, &orderSvcResData)

	helper.JoinResponse(&orderSvcResData, &orderWebResponse)

	httpUrlGetProducts := fmt.Sprintf("%s/products", controller.Config.HappyProductServiceUrl)
	svcProductsRes := helper.CreateHttpRequestService(http.MethodGet, httpUrlGetProducts, nil)

	if svcProductsRes.Status == "failed" {
		return fiberCtx.Status(fiber.StatusNotFound).JSON(svcProductsRes)
	}

	productsSvcResData := []web.ProductServiceResponse{}
	helper.GetServiceDataResponse(&svcProductsRes.Data, &productsSvcResData)

	for _, product := range productsSvcResData {
		if product.Id == orderWebResponse.IdProduct {
			helper.JoinResponse(&product, &orderWebResponse)
			break
		}
	}

	webResponse := web.WebResponse[web.OrderWebResponse]{
		Code:   fiber.StatusOK,
		Status: "success",
		Data:   orderWebResponse,
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller *OrderControllerImpl) FindAll(fiberCtx *fiber.Ctx) error {
	ordersWebResponse := []web.OrderWebResponse{}

	httpUrlGetOrders := fmt.Sprintf("%s/orders", controller.Config.HappyOrderServiceUrl)

	ordersSvcRes := helper.CreateHttpRequestService(http.MethodGet, httpUrlGetOrders, nil)

	if ordersSvcRes.Status == "failed" {
		return fiberCtx.JSON(ordersSvcRes)
	}

	ordersSvcResData := []web.OrderServiceResponse{}
	helper.GetServiceDataResponse(&ordersSvcRes.Data, &ordersSvcResData)

	httpUrlGetUsers := fmt.Sprintf("%s/users", controller.Config.HappyUserServiceUrl)

	usersSvcRes := helper.CreateHttpRequestService(http.MethodGet, httpUrlGetUsers, nil)

	usersSvcResData := []web.UserServiceResponse{}
	helper.GetServiceDataResponse(&usersSvcRes.Data, &usersSvcResData)

	for _, order := range ordersSvcResData {
		orderWebResponse := web.OrderWebResponse{}

		helper.JoinResponse(&order, &orderWebResponse)

		for _, user := range usersSvcResData {
			if order.IdUser == user.Id {
				orderWebResponse.UserName = user.UserName
				orderWebResponse.UserEmail = user.Email

				break
			}
		}

		ordersWebResponse = append(ordersWebResponse, orderWebResponse)
	}

	httpUrlGetProducts := fmt.Sprintf("%s/products", controller.Config.HappyProductServiceUrl)
	svcProductsRes := helper.CreateHttpRequestService(http.MethodGet, httpUrlGetProducts, nil)

	if svcProductsRes.Status == "failed" {
		return fiberCtx.JSON(svcProductsRes)
	}

	productsSvcResData := []web.ProductServiceResponse{}
	helper.GetServiceDataResponse(&svcProductsRes.Data, &productsSvcResData)

	for idx, order := range ordersWebResponse {
		for _, product := range productsSvcResData {
			if product.Id == order.IdProduct {
				ordersWebResponse[idx].Brand = product.Brand
				ordersWebResponse[idx].Type = product.Type
				ordersWebResponse[idx].Name = product.Name
				ordersWebResponse[idx].PriceName = product.PriceName

				break
			}
		}
	}

	webResponse := web.WebResponse[[]web.OrderWebResponse]{
		Code:   fiber.StatusOK,
		Status: "success",
		Data:   ordersWebResponse,
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(webResponse)
}
