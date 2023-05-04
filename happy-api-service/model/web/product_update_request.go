package web

type ProductUpdateRequest struct {
	Id              uint `json:"id"`
	IsOrder         bool `json:"isOrder"`
	IsOrderReject   bool `json:"isOrderReject"`
	IsPaid          bool `json:"isPaid"`
	ProductReserved uint `json:"productReserved"`
}
