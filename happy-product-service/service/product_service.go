package service

import (
	"context"
	"happy-product-service/model/web"
)

type ProductService interface {
	Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse
	FindById(ctx context.Context, categoryId uint) web.ProductResponse
	FindAll(ctx context.Context) []web.ProductResponse
}
