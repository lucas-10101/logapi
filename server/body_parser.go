package server

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/lucas-10101/logapi/server/http_utils"
)

/*
Read request Body property. Content-Type aware function

Not all content-types are supported
*/
func RequestBodyReader(request *http.Request, target *any) error {
	mediaType := http_utils.ContentType(request.Header.Get(""))
	var parseError error

	switch mediaType {
	case http_utils.MimeTypeApplicationJson:
		parseError = jsonBodyParser(request, target)
	case http_utils.MimeTypeApplicationXml:
		parseError = jsonBodyParser(request, target)
	default:
		return http_utils.NewHttpError(http.StatusUnsupportedMediaType, fmt.Sprintf("%s not supported", string(mediaType)))
	}

	if parseError != nil {
		return http_utils.NewHttpError(http.StatusBadRequest, fmt.Sprintf("cannot decode content-type %s, invalid body format", string(mediaType)))
	}
	return nil
}

func jsonBodyParser(request *http.Request, target *any) error {
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(target)
}

func xmlBodyParser(request *http.Request, target *any) error {
	decoder := xml.NewDecoder(nil)
	decoder.Strict = true
	return decoder.Decode(target)
}
