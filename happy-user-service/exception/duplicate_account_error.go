package exception

type DuplicateAccountError struct {
	Message string
}

func (err *DuplicateAccountError) Error() string {
	return err.Message
}

func NewDuplicateAccountError(err string) error {
	return &DuplicateAccountError{Message: err}
}
