package handlers

import "net/http"

func HandlerCreateNewLog(writter http.ResponseWriter, request *http.Request) {
	writter.WriteHeader(http.StatusOK)
}
