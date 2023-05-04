package controller

import (
	"happy-product-service/helper"
	"happy-product-service/model/web"
	"happy-product-service/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productUpdateRequest := web.ProductUpdateRequest{}
	helper.ReadFromRequestBody(request, &productUpdateRequest)

	productResponse := controller.ProductService.Update(request.Context(), productUpdateRequest)

	webResponse := web.WebResponse[web.ProductResponse]{
		Code:   200,
		Status: "success",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.DoPanicIfError(err)

	productResponse := controller.ProductService.FindById(request.Context(), uint(id))
	webResponse := web.WebResponse[web.ProductResponse]{
		Code:   200,
		Status: "success",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productResponses := controller.ProductService.FindAll(request.Context())
	webResponse := web.WebResponse[[]web.ProductResponse]{
		Code:   200,
		Status: "success",
		Data:   productResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
