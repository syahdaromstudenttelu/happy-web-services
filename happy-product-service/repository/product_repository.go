package repository

import (
	"context"
	"database/sql"
	"happy-product-service/model/domain"
)

type ProductRepository interface {
	Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	FindById(ctx context.Context, tx *sql.Tx, productId uint) (domain.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
}
