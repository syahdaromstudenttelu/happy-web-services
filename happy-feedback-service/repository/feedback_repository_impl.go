package repository

import (
	"context"
	"database/sql"
	"happy-feedback-service/helper"
	"happy-feedback-service/model/domain"
	"time"
)

type FeedbackRepositoryImpl struct{}

func NewFeedbackRepository() FeedbackRepository {
	return &FeedbackRepositoryImpl{}
}

func (repository *FeedbackRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, feedback domain.Feedback) domain.Feedback {
	sqlExec := "INSERT INTO feedback(id_user, id_product, feedback) VALUES (?, ?, ?)"
	result, err := tx.ExecContext(ctx, sqlExec, feedback.IdUser, feedback.IdProduct, feedback.Feedback)
	helper.DoPanicIfError(err)

	feedbackId, err := result.LastInsertId()
	helper.DoPanicIfError(err)

	feedback.Id = uint(feedbackId)
	feedback.CreatedAt = time.Now()
	return feedback
}

func (repository *FeedbackRepositoryImpl) FindByIdProduct(ctx context.Context, tx *sql.Tx, productId uint) []domain.Feedback {
	sqlQuery := "SELECT id, id_user, id_product, feedback, created_at FROM feedback WHERE id_product = ?"
	rows, err := tx.QueryContext(ctx, sqlQuery, productId)
	helper.DoPanicIfError(err)
	defer rows.Close()

	feedbacks := []domain.Feedback{}
	for rows.Next() {
		feedback := domain.Feedback{}
		rows.Scan(&feedback.Id, &feedback.IdUser, &feedback.IdProduct, &feedback.Feedback, &feedback.CreatedAt)

		if feedback.IdProduct == productId {
			feedbacks = append(feedbacks, feedback)
		}
	}

	return feedbacks
}

func (repository *FeedbackRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Feedback {
	sqlQuery := "SELECT id, id_user, id_product, feedback, created_at FROM feedback"
	rows, err := tx.QueryContext(ctx, sqlQuery)
	helper.DoPanicIfError(err)
	defer rows.Close()

	feedbacks := []domain.Feedback{}
	for rows.Next() {
		feedback := domain.Feedback{}
		rows.Scan(&feedback.Id, &feedback.IdUser, &feedback.IdProduct, &feedback.Feedback, &feedback.CreatedAt)
		feedbacks = append(feedbacks, feedback)
	}

	return feedbacks
}
