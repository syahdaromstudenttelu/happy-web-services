package web

type OrderCreateRequest struct {
	IdUser    uint `validate:"required" json:"idUser"`
	IdProduct uint `validate:"required" json:"idProduct"`
	Price     uint `validate:"required" json:"price"`
	Quantity  uint `validate:"required,min=1" json:"quantity"`
}
