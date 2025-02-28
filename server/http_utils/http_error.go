package http_utils

import (
	"fmt"
	"net/http"
)

type HttpError interface {
	Error() string
	StatusCode() int
	Send(http.ResponseWriter)
}

type httpErrorInternal struct {
	statusCode   int
	errorMessage string
}

// Create an new HttpError object
func NewHttpError(status int, message string) HttpError {
	return &httpErrorInternal{
		statusCode:   status,
		errorMessage: message,
	}
}

func (e httpErrorInternal) Error() string {
	return e.errorMessage
}

func (e httpErrorInternal) StatusCode() int {
	return e.statusCode
}

func (e httpErrorInternal) Send(writter http.ResponseWriter) {
	writter.WriteHeader(e.StatusCode())
	fmt.Fprintln(writter, e.Error())
}
