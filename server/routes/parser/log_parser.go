package parser

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/lucas-10101/logapi/data/models"
)

func ReadLogObjectFromRequest(r *http.Request) (*models.LogDocument, error) {
	data, err := io.ReadAll(r.Body)

	if err != nil {
		return nil, err
	}

	var payload models.LogDocument
	err = json.Unmarshal(data, &payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
