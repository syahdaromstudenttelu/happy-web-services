package service

import (
	"context"
	"happy-user-service/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	FindById(ctx context.Context, userId uint) web.UserResponse
	FindByUserName(ctx context.Context, username string) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}
