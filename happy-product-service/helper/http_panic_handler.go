package helper

import (
	"happy-product-service/exception"
	"happy-product-service/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func internalServerError(writer http.ResponseWriter, request *http.Request, err any) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse[any]{
		Code:   http.StatusInternalServerError,
		Status: "failed",
		Data:   err,
	}

	WriteToResponseBody(writer, webResponse)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, isError := err.(exception.NotFoundError)

	if isError {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse[any]{
			Code:   http.StatusNotFound,
			Status: "failed",
			Data:   exception.Error,
		}

		WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func reservationExceededStockError(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, isError := err.(exception.ReservationExceededError)

	if isError {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotAcceptable)

		webResponse := web.WebResponse[any]{
			Code:   http.StatusNotAcceptable,
			Status: "failed",
			Data:   exception.Error,
		}

		WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func reqBodyMalformedError(writer http.ResponseWriter, request *http.Request, err any) bool {
	_, isError := err.(exception.ReqBodyMalformedError)

	if isError {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse[any]{
			Code:   http.StatusBadRequest,
			Status: "failed",
			Data:   "request body is malformed",
		}

		WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func validationError(writer http.ResponseWriter, request *http.Request, err any) bool {
	_, isError := err.(validator.ValidationErrors)

	if isError {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse[any]{
			Code:   http.StatusBadRequest,
			Status: "failed",
			Data:   "request body value is not valid",
		}

		WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func HttpRouterPanicHandler(writer http.ResponseWriter, request *http.Request, err any) {
	if notFoundError(writer, request, err) {
		return
	}

	if reservationExceededStockError(writer, request, err) {
		return
	}

	if reqBodyMalformedError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}
