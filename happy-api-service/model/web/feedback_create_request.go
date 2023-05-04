package web

type FeedbackCreateRequest struct {
	IdUser    uint   `json:"idUser"`
	IdProduct uint   `json:"idProduct"`
	Feedback  string `json:"feedback"`
}
