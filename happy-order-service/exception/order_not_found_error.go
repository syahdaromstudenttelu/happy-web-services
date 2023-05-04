package exception

type OrderNotFoundError struct {
	Message string
}

func (error *OrderNotFoundError) Error() string {
	return error.Message
}

func NewOrderNotFoundError(err string) error {
	return &OrderNotFoundError{
		Message: err,
	}
}
