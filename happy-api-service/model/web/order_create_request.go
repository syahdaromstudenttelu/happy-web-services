package web

type OrderCreateRequest struct {
	IdUser    uint `json:"idUser"`
	IdProduct uint `json:"idProduct"`
	Price     uint `json:"price"`
	Quantity  uint `json:"quantity"`
}
