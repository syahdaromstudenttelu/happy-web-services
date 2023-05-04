package web

type UserWebResponse struct {
	Id       uint   `json:"id"`
	FullName string `json:"fullname"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}
