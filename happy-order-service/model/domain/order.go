package domain

import "time"

type Order struct {
	IdUser        uint
	IdProduct     uint
	IdOrder       string
	Price         uint
	Quantity      uint
	TotalPrice    uint
	OrderedDate   time.Time
	ExpiredDate   time.Time
	StatusPayment bool
	FeedbackDone  bool
}
