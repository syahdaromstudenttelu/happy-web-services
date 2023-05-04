package exception

type UrlParamError struct {
	Message string
}

func (err *UrlParamError) Error() string {
	return err.Message
}

func NewUrlParamError(err string) error {
	return &UrlParamError{Message: err}
}
