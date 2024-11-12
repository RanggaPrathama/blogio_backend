package responses

type ErrorResponse struct{
	ERROR  error
}


func (e ErrorResponse) Error() string{
	return e.ERROR.Error()
}

func NewErrorNotFound(err error) error {
	return &ErrorResponse{ERROR: err}
}