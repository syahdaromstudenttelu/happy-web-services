package helper

import (
	"happy-order-service/model/domain"
	"happy-order-service/model/web"
	"time"
)

func ToOrderResponse(order domain.Order) web.OrderResponse {
	return web.OrderResponse{
		IdUser:        order.IdUser,
		IdProduct:     order.IdProduct,
		IdOrder:       order.IdOrder,
		Price:         order.Price,
		Quantity:      order.Quantity,
		TotalPrice:    order.TotalPrice,
		StatusPayment: order.StatusPayment,
		OrderedDate:   order.OrderedDate.Format(time.RFC3339),
		ExpiredDate:   order.ExpiredDate.Format(time.RFC3339),
		FeedbackDone:  order.FeedbackDone,
	}
}

func ToOrdersResponse(orders []domain.Order) []web.OrderResponse {
	var ordersResponse []web.OrderResponse

	for _, order := range orders {
		ordersResponse = append(ordersResponse, ToOrderResponse(order))
	}

	return ordersResponse
}
