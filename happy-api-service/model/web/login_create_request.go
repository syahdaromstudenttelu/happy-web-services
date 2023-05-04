package web

type LoginCreateRequest struct {
	UserName string `validate:"required,alphanum,lowercase,min=1" json:"username"`
	Password string `validate:"required,min=8" json:"password"`
}
