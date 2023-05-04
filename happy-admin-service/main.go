package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"happy-admin-service/helper"
	"happy-admin-service/model/web"
	"happy-admin-service/util"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config, err := util.LoadConfig(".")
	helper.DoPanicIfError(err)

	fiberApp := fiber.New()
	fiberApp.Use(recover.New())

	allowOrigins := strings.Join(strings.Split(config.AllowOrigins, ","), ", ")

	fiberApp.Use(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowCredentials: true,
	}))

	fiberApp.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			config.AdminUname: config.AdminPass,
		},
		Unauthorized: func(c *fiber.Ctx) error {
			webResponse := web.WebResponse[string]{
				Code:   fiber.StatusUnauthorized,
				Status: "failed",
				Data:   "unauthorized",
			}

			return c.Status(fiber.StatusUnauthorized).JSON(webResponse)
		},
	}))

	// Login API
	fiberApp.Get("/login", func(c *fiber.Ctx) error {
		webResponse := web.WebResponse[web.LoginWebResponse]{
			Code:   fiber.StatusOK,
			Status: "success",
			Data: web.LoginWebResponse{
				AuthStatus:    true,
				AdminUsername: config.AdminUname,
				AdminPassword: config.AdminPass,
			},
		}

		return c.Status(fiber.StatusOK).JSON(webResponse)
	})

	// Product API
	fiberApp.Put("/products", func(fiberCtx *fiber.Ctx) error {
		reqBody := fiberCtx.Request().Body()
		httpUrl := fmt.Sprintf("%s/products", config.HappyProductServiceUrl)

		svcProductUpdateRes := helper.CreateHttpRequestService(http.MethodPut, httpUrl, bytes.NewBuffer(reqBody))

		return fiberCtx.JSON(svcProductUpdateRes)
	})

	// Order API
	fiberApp.Get("/orders", func(fiberCtx *fiber.Ctx) error {
		ordersWebResponse := []web.OrderWebResponse{}

		httpUrlGetOrders := fmt.Sprintf("%s/orders", config.HappyOrderServiceUrl)

		ordersSvcRes := helper.CreateHttpRequestService(http.MethodGet, httpUrlGetOrders, nil)

		if ordersSvcRes.Status == "failed" {
			return fiberCtx.JSON(ordersSvcRes)
		}

		ordersSvcResData := []web.OrderServiceResponse{}
		helper.GetServiceDataResponse(&ordersSvcRes.Data, &ordersSvcResData)

		httpUrlGetUsers := fmt.Sprintf("%s/users", config.HappyUserServiceUrl)

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

		httpUrlGetProducts := fmt.Sprintf("%s/products", config.HappyProductServiceUrl)
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
	})

	fiberApp.Put("/orders/statusPayment/:orderId", func(fiberCtx *fiber.Ctx) error {
		orderId := fiberCtx.Params("orderId")

		httpUrlGetOrder := fmt.Sprintf("%s/orders/orderId/%s", config.HappyOrderServiceUrl, orderId)
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

		httpUrlPutProduct := fmt.Sprintf("%s/products", config.HappyProductServiceUrl)
		helper.CreateHttpRequestService(http.MethodPut, httpUrlPutProduct, bytes.NewBuffer(productUpdReqMarshalled))

		httpUrlPutOrder := fmt.Sprintf("%s/orders/statusPayment/%s", config.HappyOrderServiceUrl, orderId)
		orderSvcPutRes := helper.CreateHttpRequestService(http.MethodPut, httpUrlPutOrder, nil)

		return fiberCtx.Status(fiber.StatusOK).JSON(orderSvcPutRes)
	})

	fiberApp.Listen(config.ServerAddr)
}
