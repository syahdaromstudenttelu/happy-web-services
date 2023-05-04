package repository

import (
	"context"
	"database/sql"
	"happy-user-service/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindById(ctx context.Context, tx *sql.Tx, userId uint) (domain.User, error)
	FindByUserName(ctx context.Context, tx *sql.Tx, username string) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
}
