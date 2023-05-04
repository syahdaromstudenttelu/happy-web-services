package test

import (
	"context"
	"fmt"
	"happy-feedback-service/app"
	"happy-feedback-service/helper"
	"happy-feedback-service/model/domain"
	"happy-feedback-service/util"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestTemp(t *testing.T) {
	config, err := util.LoadConfig("../")
	helper.DoPanicIfError(err)

	db := app.NewDB(config.DBDriver, config.DBSource)
	ctx := context.Background()

	tx, err := db.Begin()
	helper.DoPanicIfError(err)

	sqlQuery := "SELECT id, id_user, id_product, feedback, created_at FROM feedback"
	rows, err := tx.QueryContext(ctx, sqlQuery)
	helper.DoPanicIfError(err)
	defer rows.Close()

	feedbacks := []domain.Feedback{}
	for rows.Next() {
		feedback := domain.Feedback{}
		err := rows.Scan(&feedback.Id, &feedback.IdUser, &feedback.IdProduct, &feedback.Feedback, &feedback.CreatedAt)
		helper.DoPanicIfError(err)
		feedbacks = append(feedbacks, feedback)
	}

	fmt.Println(feedbacks[0].CreatedAt.Format(time.RFC3339))
}
