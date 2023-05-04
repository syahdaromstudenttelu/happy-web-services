package exception

type ReqBodyMalformedError struct {
	Message string
}

func (err *ReqBodyMalformedError) Error() string {
	return err.Message
}

func NewReqBodyMalformedError(err string) error {
	return &ReqBodyMalformedError{Message: err}
}
