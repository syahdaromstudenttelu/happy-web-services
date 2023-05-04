package web

type FeedbackUserRequest struct {
	IdProduct uint   `json:"idProduct"`
	IdOrder   string `json:"idOrder"`
	Feedback  string `json:"feedback"`
}
