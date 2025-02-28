package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type LogModel struct {
	Application      string
	ApplicationRoute string
	RequestMethod    string
	StatusCode       int
	IsServerError    bool
	IsClientError    bool
	ErrorMessage     string
}

func (model *LogModel) Validate() bool {
	if model.Application == "" {
		return false
	}
	if model.ApplicationRoute == "" {
		return false
	}
	if model.RequestMethod == "" {
		return false
	}
	if model.StatusCode < 0 {
		return false
	}
	if model.IsServerError && model.IsClientError {
		return false
	}
	if model.IsServerError && model.ErrorMessage == "" {
		return false
	}
	return true
}

func (model *LogModel) ToBsonD() bson.D {
	return bson.D{
		{
			Key:   "Timestamp",
			Value: time.Now(),
		},
		{
			Key:   "Application",
			Value: model.Application,
		},
		{
			Key:   "ApplicationRoute",
			Value: model.ApplicationRoute,
		},
		{
			Key:   "RequestMethod",
			Value: model.RequestMethod,
		},
		{
			Key:   "StatusCode",
			Value: model.StatusCode,
		},
		{
			Key:   "IsServerError",
			Value: model.IsServerError,
		},
		{
			Key:   "IsClientError",
			Value: model.IsClientError,
		},
		{
			Key:   "ErrorMessage",
			Value: model.ErrorMessage,
		},
	}
}
