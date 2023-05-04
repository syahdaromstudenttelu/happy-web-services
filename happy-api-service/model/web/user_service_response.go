package web

type UserServiceResponse struct {
	Id       uint   `json:"id"`
	FullName string `json:"fullname"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
