package web

type FeedbackResponse struct {
	Id        uint   `json:"id"`
	IdUser    uint   `json:"idUser"`
	IdProduct uint   `json:"idProduct"`
	Feedback  string `json:"feedback"`
	CreatedAt string `json:"createdAt"`
}
