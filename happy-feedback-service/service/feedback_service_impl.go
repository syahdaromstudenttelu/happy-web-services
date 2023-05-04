package service

import (
	"context"
	"database/sql"
	"happy-feedback-service/exception"
	"happy-feedback-service/helper"
	"happy-feedback-service/model/domain"
	"happy-feedback-service/model/web"
	"happy-feedback-service/repository"

	"github.com/go-playground/validator/v10"
)

type FeedbackServiceImpl struct {
	FeedbackRepository repository.FeedbackRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewFeedbackService(feedbackRepository repository.FeedbackRepository, db *sql.DB, validate *validator.Validate) FeedbackService {
	return &FeedbackServiceImpl{
		FeedbackRepository: feedbackRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *FeedbackServiceImpl) Create(ctx context.Context, request web.FeedbackCreateRequest) web.FeedbackResponse {
	err := service.Validate.Struct(request)

	if err != nil {
		helper.DoPanicIfError(exception.NewReqBodyMalformedError("request body is not valid"))
	}

	tx, err := service.DB.Begin()
	helper.DoPanicIfError(err)
	defer helper.CommitOrRollback(tx)

	feedback := domain.Feedback{
		IdUser:    request.IdUser,
		IdProduct: request.IdProduct,
		Feedback:  request.Feedback,
	}

	feedback = service.FeedbackRepository.Save(ctx, tx, feedback)
	return helper.ToFeedbackResponse(feedback)
}

func (service *FeedbackServiceImpl) FindByIdProduct(ctx context.Context, productId uint) []web.FeedbackResponse {
	tx, err := service.DB.Begin()
	helper.DoPanicIfError(err)
	defer helper.CommitOrRollback(tx)

	feedbacks := service.FeedbackRepository.FindByIdProduct(ctx, tx, productId)

	return helper.ToFeedbacksResponse(feedbacks)
}

func (service *FeedbackServiceImpl) FindAll(ctx context.Context) []web.FeedbackResponse {
	tx, err := service.DB.Begin()
	helper.DoPanicIfError(err)
	defer helper.CommitOrRollback(tx)

	feedbacks := service.FeedbackRepository.FindAll(ctx, tx)

	return helper.ToFeedbacksResponse(feedbacks)
}
