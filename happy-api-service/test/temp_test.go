package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"happy-api-service/helper"
	"happy-api-service/model/web"
	"happy-api-service/util"
	"net/http"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestFiber(t *testing.T) {
	fiberApp := fiber.New()

	fiberApp.Put("/test", func(c *fiber.Ctx) error {
		bodyRequest := fiber.Map{}
		err := json.Unmarshal(c.Request().Body(), &bodyRequest)
		helper.DoPanicIfError(err)

		fmt.Println(bodyRequest)

		return c.JSON(fiber.Map{
			"code": fiber.StatusOK,
			"data": bodyRequest,
		})
	})

	err := fiberApp.Listen(":3000")
	helper.DoPanicIfError(err)
}

func TestPut(t *testing.T) {
	url := "http://localhost:3000/test"
	reqBody := web.ProductUpdateRequest{
		Id:              1,
		IsOrder:         true,
		ProductReserved: 500,
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	helper.DoPanicIfError(err)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(reqBodyBytes))

	helper.DoPanicIfError(err)

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	helper.DoPanicIfError(err)
	defer response.Body.Close()

	responseMap := fiber.Map{}
	json.NewDecoder(response.Body).Decode(&responseMap)

	fmt.Println(responseMap)
}

func TestUpdateProduct(t *testing.T) {
	// config, err := util.LoadConfig("../")
	// helper.DoPanicIfError(err)

	// httpUrl := fmt.Sprintf("%s/products/%d", config.HappyProductServiceUrl, 2)
	// request, err := http.NewRequest(http.MethodPut, httpUrl, )
	// helper.DoPanicIfError(err)

	// product := domain.HttpResponse[web.ProductResponse]{}
	// err = json.NewDecoder(response.Body).Decode(&product)
	// helper.DoPanicIfError(err)

	// fmt.Println(product.Code, product.Status, product.Data)
}

func TestGetProduct(t *testing.T) {
	config, err := util.LoadConfig("../")
	helper.DoPanicIfError(err)

	httpUrl := fmt.Sprintf("%s/products/%d", config.HappyProductServiceUrl, 2)
	response, err := http.Get(httpUrl)
	helper.DoPanicIfError(err)

	product := web.WebResponse[web.ProductWebResponse]{}
	err = json.NewDecoder(response.Body).Decode(&product)
	helper.DoPanicIfError(err)

	fmt.Println(product.Code, product.Status, product.Data)
}

func TestGetProducts(t *testing.T) {
	config, err := util.LoadConfig("../")
	helper.DoPanicIfError(err)

	httpUrl := fmt.Sprintf("%s/products", config.HappyProductServiceUrl)
	response, err := http.Get(httpUrl)
	helper.DoPanicIfError(err)

	products := web.WebResponse[[]web.ProductWebResponse]{}
	err = json.NewDecoder(response.Body).Decode(&products)
	helper.DoPanicIfError(err)

	fmt.Println(products.Code, products.Status)

	for _, product := range products.Data {
		fmt.Println(product)
	}
}

func TestGetUsers(t *testing.T) {
	config, err := util.LoadConfig("../")
	helper.DoPanicIfError(err)

	httpUrl := fmt.Sprintf("%s/users", config.HappyUserServiceUrl)
	response, err := http.Get(httpUrl)
	helper.DoPanicIfError(err)

	users := web.WebResponse[any]{}
	err = json.NewDecoder(response.Body).Decode(&users)
	helper.DoPanicIfError(err)

	fmt.Println(users.Code, users.Status)

	for _, product := range users.Data.([]any) {
		fmt.Println(product.(map[string]any)["id"])
	}
}

func TestConfig(t *testing.T) {
	config, err := util.LoadConfig("../")
	helper.DoPanicIfError(err)

	fmt.Println(config)
}

func TestTemp(t *testing.T) {
	config, err := util.LoadConfig("../")
	helper.DoPanicIfError(err)

	allowOrigins := strings.Join(strings.Split(config.AllowOrigins, ","), ", ")

	fmt.Println(allowOrigins)
}
