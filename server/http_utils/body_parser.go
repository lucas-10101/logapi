package http_utils

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

/*
Read request Body property. Content-Type aware function

Not all content-types are supported
*/
func RequestBodyParser(request *http.Request, target any) HttpError {

	switch ContentType(request.Header.Get(string(ContentTypeHeader))) {
	case MimeTypeApplicationJson:
		return jsonBodyParser(request, target)
	case MimeTypeApplicationXml:
		return xmlBodyParser(request, target)
	default:
		return NewHttpError(http.StatusUnsupportedMediaType, fmt.Sprintf("%s not supported", request.Header.Get(string(ContentTypeHeader))))
	}
}

func jsonBodyParser(request *http.Request, target any) HttpError {
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(target)

	if err != nil {
		if err == io.EOF {
			return NewHttpError(http.StatusBadRequest, err.Error())
		}
		return NewHttpError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func xmlBodyParser(request *http.Request, target any) HttpError {
	decoder := xml.NewDecoder(request.Body)
	decoder.Strict = true
	err := decoder.Decode(target)

	if err != nil {
		if err == io.EOF {
			return NewHttpError(http.StatusBadRequest, err.Error())
		}
		return NewHttpError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
