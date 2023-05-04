package web

type OrderWebResponse struct {
	UserName      string `json:"username"`
	UserEmail     string `json:"userEmail"`
	IdOrder       string `json:"idOrder"`
	IdProduct     uint   `json:"idProduct"`
	Brand         string `json:"brand"`
	Type          string `json:"type"`
	Name          string `json:"name"`
	PriceName     string `json:"priceName"`
	Price         uint   `json:"price"`
	Quantity      uint   `json:"quantity"`
	TotalPrice    uint   `json:"totalPrice"`
	OrderedDate   string `json:"orderedDate"`
	ExpiredDate   string `json:"expiredDate"`
	StatusPayment bool   `json:"statusPayment"`
	FeedbackDone  bool   `json:"feedbackDone"`
}
