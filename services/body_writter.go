package services

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/lucas-10101/logapi/web/webutils"
)

func BodyWritter(writter http.ResponseWriter, request *http.Request, src any) webutils.HttpError {
	headerValue, exists := request.Header[string(webutils.HeaderContentType)]
	if !exists || len(headerValue) < 1 {
		return webutils.NewHttpError(http.StatusUnsupportedMediaType, "unknown media type")
	}

	switch webutils.MimeType(headerValue[0]) {
	case webutils.MimeTypeApplicationJson:
		writter.Header().Add(string(webutils.HeaderContentType), string(webutils.MimeTypeApplicationJson))
		if err := jsonBodyWritter(writter, src); err != nil {
			return webutils.NewHttpError(http.StatusBadRequest, "an error has occoured when writting json response")
		}

	case webutils.MimeTypeApplicationXml:
		writter.Header().Add(string(webutils.HeaderContentType), string(webutils.MimeTypeApplicationXml))
		if err := xmlBodyWritter(writter, src); err != nil {
			return webutils.NewHttpError(http.StatusBadRequest, "an error has occoured when writting xml response")
		}
	default:
		return webutils.NewHttpError(http.StatusUnsupportedMediaType, "unsupported media type")
	}

	return nil
}

func jsonBodyWritter(writter http.ResponseWriter, data any) error {
	encoder := json.NewEncoder(writter)
	encoder.SetIndent("", "")
	return encoder.Encode(data)
}

func xmlBodyWritter(writter http.ResponseWriter, data any) error {
	encoder := xml.NewEncoder(writter)
	return encoder.Encode(data)
}
