package webutils

type HttpError interface {
	Error() string
	StatusCode() int
}

type httpErrorInternal struct {
	errorMessage string
	statusCode   int
}

func (err *httpErrorInternal) Error() string {
	return err.errorMessage
}

func (err *httpErrorInternal) StatusCode() int {
	return err.statusCode
}

func NewHttpError(statusCode int, errorMessage string) HttpError {
	return &httpErrorInternal{
		statusCode:   statusCode,
		errorMessage: errorMessage,
	}
}
