package service

import (
	"context"
	"happy-order-service/model/web"
)

type OrderService interface {
	Create(ctx context.Context, request web.OrderCreateRequest) web.OrderResponse
	UpdateStatusPayment(ctx context.Context, orderId string)
	UpdateFeedbackDone(ctx context.Context, orderId string)
	Delete(ctx context.Context, userId uint, orderId string)
	FindByIdOrder(ctx context.Context, orderId string) web.OrderResponse
	FindByIdUserAndIdOrder(ctx context.Context, userId uint, orderId string) web.OrderResponse
	FindAllByIdUser(ctx context.Context, userId uint) []web.OrderResponse
	FindAll(ctx context.Context) []web.OrderResponse
}
