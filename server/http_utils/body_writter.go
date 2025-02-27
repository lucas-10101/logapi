package http_utils

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

func ResponseBodyWritter(contentType ContentType, writter http.ResponseWriter, data any) HttpError {
	switch contentType {
	case MimeTypeApplicationJson:
		return jsonBodyWritter(writter, data)
	case MimeTypeApplicationXml:
		return xmlBodyWritter(writter, data)
	default:
		return NewHttpError(http.StatusUnsupportedMediaType, fmt.Sprintf("%s not supported", string(contentType)))
	}
}

func jsonBodyWritter(writter http.ResponseWriter, data any) HttpError {

	var jsonData []byte
	jsonData, err := json.Marshal(data)

	if err != nil {
		return NewHttpError(http.StatusInternalServerError, err.Error())
	}

	writter.Header().Add(string(ContentTypeHeader), string(MimeTypeApplicationJson))
	writter.Write(jsonData)

	return nil
}

func xmlBodyWritter(writter http.ResponseWriter, data any) HttpError {

	var xmlData []byte
	xmlData, err := xml.Marshal(data)

	if err != nil {
		return NewHttpError(http.StatusInternalServerError, err.Error())
	}

	writter.Header().Add(string(ContentTypeHeader), string(MimeTypeApplicationXml))
	writter.Write(xmlData)

	return nil
}
