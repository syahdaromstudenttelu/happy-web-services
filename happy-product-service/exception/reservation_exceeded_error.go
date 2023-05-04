package exception

type ReservationExceededError struct {
	Error string
}

func NewReservationExceededError(err string) ReservationExceededError {
	return ReservationExceededError{Error: err}
}
