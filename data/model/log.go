package model

import "time"

// Represents basic structure of log trace and metrics data
type LogDocument struct {
	Timestamp         time.Time
	ApplicationSource LogDocumentApplicationSource
	RequestInfo       LogDocumentRequestInfo
	ErrorInfo         LogDocumentErrorInfo
}

type LogDocumentApplicationSource struct {
	ApplicationName  string
	ApplicationRoute string
}

type LogDocumentRequestInfo struct {
	RequestStatusCode int
	RequestMethod     string
	RequestDuration   time.Duration
}

type LogDocumentErrorInfo struct {
	StackTraceMessage string
}
