package repository

import (
	"context"
	"database/sql"
	"errors"
	"happy-user-service/exception"
	"happy-user-service/helper"
	"happy-user-service/model/domain"
	"strings"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	sqlExect := "INSERT INTO user(fullname, username, email, password) VALUES (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, sqlExect, user.FullName, user.UserName, user.Email, user.Password)

	if err != nil {
		errMsg := err.Error()

		if strings.Contains(errMsg, "user.username_unique") {
			panic(exception.NewDuplicateAccountError("username_duplicate"))
		}

		if strings.Contains(errMsg, "user.email_unique") {
			panic(exception.NewDuplicateAccountError("email_duplicate"))
		}
	}

	userId, err := result.LastInsertId()
	helper.DoPanicIfError(err)

	user.Id = uint(userId)
	return user
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId uint) (domain.User, error) {
	sqlQuery := "SELECT id, fullname, username, email, password FROM user WHERE id = ?"
	rows, err := tx.QueryContext(ctx, sqlQuery, userId)
	helper.DoPanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.FullName, &user.UserName, &user.Email, &user.Password)
		helper.DoPanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

func (repository *UserRepositoryImpl) FindByUserName(ctx context.Context, tx *sql.Tx, username string) (domain.User, error) {
	sqlQuery := "SELECT id, fullname, username, email, password FROM user WHERE username = ?"
	rows, err := tx.QueryContext(ctx, sqlQuery, username)
	helper.DoPanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.FullName, &user.UserName, &user.Email, &user.Password)
		helper.DoPanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	sqlQuery := "SELECT id, fullname, username, email, password FROM user"
	rows, err := tx.QueryContext(ctx, sqlQuery)
	helper.DoPanicIfError(err)
	defer rows.Close()

	users := []domain.User{}
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.FullName, &user.UserName, &user.Email, &user.Password)
		helper.DoPanicIfError(err)
		users = append(users, user)
	}

	return users
}
