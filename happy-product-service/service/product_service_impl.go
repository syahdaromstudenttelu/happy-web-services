package service

import (
	"context"
	"database/sql"
	"happy-product-service/exception"
	"happy-product-service/helper"
	"happy-product-service/model/web"
	"happy-product-service/repository"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, db *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                db,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.DoPanicIfError(err)

	tx, err := service.DB.Begin()
	helper.DoPanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, request.Id)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if request.IsOrder {
		product.Reservation = product.Reservation + request.ProductReserved
	}

	if request.IsOrderReject {
		product.Reservation = product.Reservation - request.ProductReserved
	}

	if request.IsPaid {
		product.ProductStock = product.ProductStock - request.ProductReserved
		product.Reservation = product.Reservation - request.ProductReserved
	}

	product = service.ProductRepository.Update(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId uint) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.DoPanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.DoPanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.FindAll(ctx, tx)

	return helper.ToProductsResponse(products)
}
