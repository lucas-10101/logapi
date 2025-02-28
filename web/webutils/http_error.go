package webutils

import "net/http"

type HttpError interface {
	Error() string
	StatusCode() int
	ErrorBytes() []byte
	Write(writter http.ResponseWriter)
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

func (err *httpErrorInternal) ErrorBytes() []byte {
	return []byte(err.Error())
}

func (err *httpErrorInternal) Write(writter http.ResponseWriter) {
	writter.WriteHeader(err.StatusCode())
	writter.Write(err.ErrorBytes())
}

func NewHttpError(statusCode int, errorMessage string) HttpError {
	return &httpErrorInternal{
		statusCode:   statusCode,
		errorMessage: errorMessage,
	}
}
