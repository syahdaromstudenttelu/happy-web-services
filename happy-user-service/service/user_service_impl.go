package service

import (
	"context"
	"database/sql"
	"happy-user-service/exception"
	"happy-user-service/helper"
	"happy-user-service/model/domain"
	"happy-user-service/model/web"
	"happy-user-service/repository"
	"strings"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)

	if err != nil {
		helper.DoPanicIfError(err)
	}

	tx, err := service.DB.Begin()
	helper.DoPanicIfError(err)
	defer helper.CommitOrRollback(tx)

	username := strings.ToLower(strings.Join(strings.Split(request.UserName, " "), ""))
	email := strings.ToLower(strings.Join(strings.Split(request.Email, " "), ""))
	hashedPassword := helper.BcryptPassword(request.Password)

	user := domain.User{
		FullName: request.FullName,
		UserName: username,
		Email:    email,
		Password: hashedPassword,
	}

	user = service.UserRepository.Save(ctx, tx, user)
	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId uint) web.UserResponse {
	tx, errUserNotFound := service.DB.Begin()
	helper.DoPanicIfError(errUserNotFound)
	defer helper.CommitOrRollback(tx)

	user, errUserNotFound := service.UserRepository.FindById(ctx, tx, userId)

	if errUserNotFound != nil {
		panic(exception.NewUserNotFoundError(errUserNotFound.Error()))
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindByUserName(ctx context.Context, username string) web.UserResponse {
	tx, errUserNotFound := service.DB.Begin()
	helper.DoPanicIfError(errUserNotFound)
	defer helper.CommitOrRollback(tx)

	user, errUserNotFound := service.UserRepository.FindByUserName(ctx, tx, username)

	if errUserNotFound != nil {
		panic(exception.NewUserNotFoundError(errUserNotFound.Error()))
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.DoPanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)

	return helper.ToUsersResponse(users)
}
