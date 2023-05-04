package web

type FeedbackWebResponse struct {
	Id        uint   `json:"id"`
	FullName  string `json:"fullName"`
	Feedback  string `json:"feedback"`
	CreatedAt string `json:"createdAt"`
}
