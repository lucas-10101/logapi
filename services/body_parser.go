package services

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/lucas-10101/logapi/web/webutils"
)

func BodyParser(request *http.Request, dest any) webutils.HttpError {

	headerValue, exists := request.Header[string(webutils.HeaderContentType)]
	if !exists || len(headerValue) < 1 {
		return webutils.NewHttpError(http.StatusUnsupportedMediaType, "unknown media type")
	}

	switch headerValue[0] {
	case string(webutils.MimeTypeApplicationJson):
		if err := jsonBodyParser(request, dest); err != nil {
			return webutils.NewHttpError(http.StatusBadRequest, "malformed json body")
		}
	case string(webutils.MimeTypeApplicationXml):
		if err := xmlBodyParser(request, dest); err != nil {
			return webutils.NewHttpError(http.StatusBadRequest, "malformed xml body")
		}
	default:
		return webutils.NewHttpError(http.StatusUnsupportedMediaType, "unsupported media type")
	}

	return nil
}

func jsonBodyParser(request *http.Request, dest any) error {
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(dest)
}

func xmlBodyParser(request *http.Request, dest any) error {
	decoder := xml.NewDecoder(request.Body)
	return decoder.Decode(dest)
}
