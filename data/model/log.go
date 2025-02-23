package model

import (
	"time"
)

// Represents basic structure of log trace and metrics data
type LogDocument struct {
	Timestamp         time.Time
	ApplicationSource LogDocumentApplicationSource
	RequestInfo       LogDocumentRequestInfo
	ErrorInfo         LogDocumentErrorInfo
}

func (src LogDocument) ToDocument() Document {
	dst := Document{
		{
			Field: "Timestamp",
			Value: src.Timestamp,
		},
		{
			Field: "ApplicationSource",
			Value: src.ApplicationSource.ToDocument(),
		},
		{
			Field: "RequestInfo",
			Value: src.RequestInfo.ToDocument(),
		},
		{
			Field: "ErrorInfo",
			Value: src.ErrorInfo.ToDocument(),
		},
	}

	return dst
}

type LogDocumentApplicationSource struct {
	ApplicationName  string
	ApplicationRoute string
}

func (src LogDocumentApplicationSource) ToDocument() Document {
	dst := Document{
		{
			Field: "ApplicationName",
			Value: src.ApplicationName,
		},
		{
			Field: "ApplicationRoute",
			Value: src.ApplicationRoute,
		},
	}

	return dst
}

type LogDocumentRequestInfo struct {
	RequestStatusCode int
	RequestMethod     string
	RequestDuration   *time.Duration
}

func (src LogDocumentRequestInfo) ToDocument() Document {
	dst := Document{
		{
			Field: "RequestStatusCode",
			Value: src.RequestStatusCode,
		},
		{
			Field: "RequestMethod",
			Value: src.RequestMethod,
		},
		{
			Field: "RequestDuration",
			Value: src.RequestDuration,
		},
	}

	return dst
}

type LogDocumentErrorInfo struct {
	StackTraceMessage *string
}

func (src LogDocumentErrorInfo) ToDocument() Document {
	dst := Document{
		{
			Field: "StackTraceMessage",
			Value: src.StackTraceMessage,
		},
	}

	return dst
}
