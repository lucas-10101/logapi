package http_utils

type httpError struct {
	statusCode   int
	errorMessage string
}

// Create an new HttpError object
func NewHttpError(status int, message string) *httpError {
	return &httpError{
		statusCode:   status,
		errorMessage: message,
	}
}

func (e *httpError) Error() string {
	return e.errorMessage
}

func (e *httpError) StatusCode() int {
	return e.statusCode
}
