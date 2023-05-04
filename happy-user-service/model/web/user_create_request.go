package web

type UserCreateRequest struct {
	FullName string `validate:"required,min=1" json:"fullname"`
	UserName string `validate:"required,alphanum,lowercase,min=1" json:"username"`
	Email    string `validate:"required,email,min=1" json:"email"`
	Password string `validate:"required,min=8" json:"password"`
}
