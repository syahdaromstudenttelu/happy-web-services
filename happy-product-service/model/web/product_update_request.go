package web

type ProductUpdateRequest struct {
	Id              uint `validate:"required"`
	IsOrder         bool `validate:"required_if=IsOrderReject false IsPaid false"`
	IsOrderReject   bool `validate:"required_if=IsOrder false IsPaid false"`
	IsPaid          bool `validate:"required_if=IsOrder false IsOrderReject false"`
	ProductReserved uint `validate:"required,min=1"`
}
