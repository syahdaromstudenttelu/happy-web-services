package web

type ServiceResponse struct {
	Code   uint   `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}
