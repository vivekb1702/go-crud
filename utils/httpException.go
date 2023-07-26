package utils

type CustomError struct {
	Code    int
	Message string
}

func (ce CustomError) Error() string {
	return ce.Message
}
