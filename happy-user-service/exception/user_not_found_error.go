package exception

type UserNotFoundError struct {
	Message string
}

func (error *UserNotFoundError) Error() string {
	return error.Message
}

func NewUserNotFoundError(err string) error {
	return &UserNotFoundError{
		Message: err,
	}
}
