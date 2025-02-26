package http_utils

type HttpHeader string
type ContentType string

const (
	ContentTypeHeader = HttpHeader("Content-Type")
)

const (
	MimeTypeApplicationJson = ContentType("application/json")
	MimeTypeApplicationXml  = ContentType("application/xml")
)
