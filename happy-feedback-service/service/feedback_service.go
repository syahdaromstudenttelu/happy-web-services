package service

import (
	"context"
	"happy-feedback-service/model/web"
)

type FeedbackService interface {
	Create(ctx context.Context, request web.FeedbackCreateRequest) web.FeedbackResponse
	FindByIdProduct(ctx context.Context, productId uint) []web.FeedbackResponse
	FindAll(ctx context.Context) []web.FeedbackResponse
}
