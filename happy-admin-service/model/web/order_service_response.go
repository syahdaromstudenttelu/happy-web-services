package web

type OrderServiceResponse struct {
	IdUser        uint   `json:"idUser"`
	IdProduct     uint   `json:"idProduct"`
	IdOrder       string `json:"idOrder"`
	Price         uint   `json:"price"`
	Quantity      uint   `json:"quantity"`
	TotalPrice    uint   `json:"totalPrice"`
	OrderedDate   string `json:"orderedDate"`
	ExpiredDate   string `json:"expiredDate"`
	StatusPayment bool   `json:"statusPayment"`
	FeedbackDone  bool   `json:"feedbackDone"`
}
