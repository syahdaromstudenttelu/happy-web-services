package domain

import "time"

type Feedback struct {
	Id        uint
	IdUser    uint
	IdProduct uint
	Feedback  string
	CreatedAt time.Time
}
