package helper

import (
	"happy-product-service/model/domain"
	"happy-product-service/model/web"
)

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:           product.Id,
		Brand:        product.Brand,
		Type:         product.Type,
		Name:         product.Name,
		PriceName:    product.PriceName,
		ProductPrice: product.ProductPrice,
		ProductStock: product.ProductStock,
		Reservation:  product.Reservation,
	}
}

func ToProductsResponse(products []domain.Product) []web.ProductResponse {
	productsResponse := []web.ProductResponse{}

	for _, product := range products {
		productsResponse = append(productsResponse, ToProductResponse(product))
	}

	return productsResponse
}
