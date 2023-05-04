package service

import (
	"context"
	"database/sql"
	"happy-order-service/exception"
	"happy-order-service/helper"
	"happy-order-service/model/domain"
	"happy-order-service/model/web"
	"happy-order-service/repository"

	"github.com/go-playground/validator/v10"
)

type OrderServiceImpl struct {
	OrderRepository repository.OrderRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewOrderService(orderRepository repository.OrderRepository, db *sql.DB, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrderRepository: orderRepository,
		DB:              db,
		Validate:        validate,
	}
}

func (service *OrderServiceImpl) Create(ctx context.Context, request web.OrderCreateRequest) web.OrderResponse {
	err := service.Validate.Struct(request)

	if err != nil {
		helper.DoPanicIfError(exception.NewReqBodyMalformedError("request body is not valid"))
	}

	tx, err := service.DB.Begin()
	helper.DoPanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order := domain.Order{
		IdUser:    request.IdUser,
		IdProduct: request.IdProduct,
		Price:     request.Price,
		Quantity:  request.Quantity,
	}

	order = service.OrderRepository.Save(ctx, tx, order)
	return helper.ToOrderResponse(order)
}

func (service *OrderServiceImpl) UpdateStatusPayment(ctx context.Context, orderId string) {
	tx, err := service.DB.Begin()
	helper.DoPanicIfError(err)
	defer helper.CommitOrRollback(tx)

	service.OrderRepository.UpdateStatusPayment(ctx, tx, orderId)
}

func (service *OrderServiceImpl) UpdateFeedbackDone(ctx context.Context, orderId string) {
	tx, err := service.DB.Begin()
	helper.DoPanicIfError(err)
	defer helper.CommitOrRollback(tx)

	service.OrderRepository.UpdateFeedbackDone(ctx, tx, orderId)
}

func (service *OrderServiceImpl) Delete(ctx context.Context, userId uint, orderId string) {
	tx, err := service.DB.Begin()
	helper.DoPanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order, err := service.OrderRepository.FindByIdUserAndIdOrder(ctx, tx, userId, orderId)

	if err != nil {
		panic(exception.NewOrderNotFoundError(err.Error()))
	}

	service.OrderRepository.Delete(ctx, tx, order.IdOrder)
}

func (service *OrderServiceImpl) FindByIdOrder(ctx context.Context, orderId string) web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.DoPanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order, err := service.OrderRepository.FindByIdOrder(ctx, tx, orderId)

	if err != nil {
		panic(exception.NewOrderNotFoundError(err.Error()))
	}

	return helper.ToOrderResponse(order)
}

func (service *OrderServiceImpl) FindByIdUserAndIdOrder(ctx context.Context, userId uint, orderId string) web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.DoPanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order, err := service.OrderRepository.FindByIdUserAndIdOrder(ctx, tx, userId, orderId)

	if err != nil {
		panic(exception.NewOrderNotFoundError(err.Error()))
	}

	return helper.ToOrderResponse(order)
}

func (service *OrderServiceImpl) FindAllByIdUser(ctx context.Context, userId uint) []web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.DoPanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orders := service.OrderRepository.FindAllByIdUser(ctx, tx, userId)

	return helper.ToOrdersResponse(orders)
}

func (service *OrderServiceImpl) FindAll(ctx context.Context) []web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.DoPanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orders := service.OrderRepository.FindAll(ctx, tx)
	return helper.ToOrdersResponse(orders)
}
