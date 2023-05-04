package helper

import (
	"happy-feedback-service/model/domain"
	"happy-feedback-service/model/web"
	"time"
)

func ToFeedbackResponse(feedback domain.Feedback) web.FeedbackResponse {
	return web.FeedbackResponse{
		Id:        feedback.Id,
		IdUser:    feedback.IdUser,
		IdProduct: feedback.IdProduct,
		Feedback:  feedback.Feedback,
		CreatedAt: feedback.CreatedAt.Format(time.RFC3339),
	}
}

func ToFeedbacksResponse(feedbacks []domain.Feedback) []web.FeedbackResponse {
	var feedbacksResponse []web.FeedbackResponse

	for _, feedback := range feedbacks {
		feedbacksResponse = append(feedbacksResponse, ToFeedbackResponse(feedback))
	}

	return feedbacksResponse
}
