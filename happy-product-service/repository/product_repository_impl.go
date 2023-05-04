package repository

import (
	"context"
	"database/sql"
	"errors"
	"happy-product-service/exception"
	"happy-product-service/helper"
	"happy-product-service/model/domain"
	"strings"
)

type ProductRepositoryImpl struct{}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	defer func() {
		err := recover()
		if err != nil {
			errorMsgCaused := err.(error).Error()

			if strings.Contains(errorMsgCaused, "reservation_check") {
				panic(exception.NewReservationExceededError("reservation exceeded stock"))
			}
		}
	}()

	SQL := "update product set product_stock = ?, reservation = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.ProductStock, product.Reservation, product.Id)
	helper.DoPanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId uint) (domain.Product, error) {
	SQL := "select id, brand, type, name, price_name, product_price, product_stock, reservation from product where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.DoPanicIfError(err)
	defer rows.Close()

	product := domain.Product{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Brand, &product.Type, &product.Name, &product.PriceName, &product.ProductPrice, &product.ProductStock, &product.Reservation)
		helper.DoPanicIfError(err)
		return product, nil
	} else {
		return product, errors.New("product is not found")
	}
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := "select id, brand, type, name, price_name, product_price, product_stock, reservation from product"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.DoPanicIfError(err)
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.Id, &product.Brand, &product.Type, &product.Name, &product.PriceName, &product.ProductPrice, &product.ProductStock, &product.Reservation)
		helper.DoPanicIfError(err)
		products = append(products, product)
	}

	return products
}
