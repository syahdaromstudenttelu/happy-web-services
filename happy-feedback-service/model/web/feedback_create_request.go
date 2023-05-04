package web

type FeedbackCreateRequest struct {
	IdUser    uint   `validate:"required" json:"idUser"`
	IdProduct uint   `validate:"required" json:"idProduct"`
	Feedback  string `validate:"required,min=1" json:"feedback"`
}
