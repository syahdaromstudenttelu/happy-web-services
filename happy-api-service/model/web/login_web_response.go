package web

type LoginWebResponse struct {
	Id       uint   `json:"id"`
	FullName string `json:"fullname"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}
