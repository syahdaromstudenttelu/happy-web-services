package web

type ProductWebResponse struct {
	Id           uint                  `json:"id"`
	Brand        string                `json:"brand"`
	Type         string                `json:"type"`
	Name         string                `json:"name"`
	PriceName    string                `json:"priceName"`
	ProductStock uint                  `json:"productStock"`
	ProductPrice uint                  `json:"productPrice"`
	Reservation  uint                  `json:"reservation"`
	Feedbacks    []FeedbackWebResponse `json:"feedbacks"`
}
