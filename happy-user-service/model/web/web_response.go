package web

type WebResponse[T any] struct {
	Code   uint   `json:"code"`
	Status string `json:"status"`
	Data   T      `json:"data"`
}
