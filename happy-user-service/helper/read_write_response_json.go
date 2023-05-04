package helper

import (
	"encoding/json"
	"happy-user-service/exception"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func ReadFromRequestBody[T any](request *fasthttp.Request, requestTemplate T) {
	err := json.Unmarshal(request.Body(), &requestTemplate)

	if err != nil {
		panic(exception.NewReqBodyMalformedError(err.Error()))
	}
}

func WriteToResponseBody[T any](fiberCtx *fiber.Ctx, statusCode int, responseBody T) error {
	fiberCtx.Response().Header.Add("Content-Type", "application/json")
	return fiberCtx.Status(statusCode).JSON(responseBody)
}
