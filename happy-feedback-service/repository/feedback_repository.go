package repository

import (
	"context"
	"database/sql"
	"happy-feedback-service/model/domain"
)

type FeedbackRepository interface {
	Save(ctx context.Context, tx *sql.Tx, feedback domain.Feedback) domain.Feedback
	FindByIdProduct(ctx context.Context, tx *sql.Tx, productId uint) []domain.Feedback
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Feedback
}
