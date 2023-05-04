package helper

import (
	"happy-user-service/model/domain"
	"happy-user-service/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:       user.Id,
		FullName: user.FullName,
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ToUsersResponse(users []domain.User) []web.UserResponse {
	var usersResponse []web.UserResponse

	for _, user := range users {
		usersResponse = append(usersResponse, ToUserResponse(user))
	}

	return usersResponse
}
