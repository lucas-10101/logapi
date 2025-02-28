package models

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/lucas-10101/logapi/settings"
)

// Represents an pagination request, pageNumber is 0...n
type PageRequest struct {
	PageNumber int64
	PageSize   int64
}

// Read url encoded queries (a=1&b=2) then load into page request object
func (page *PageRequest) LoadFromUrlQuery(encodedQuery string) error {

	page.PageNumber = 0
	page.PageSize = 10

	values, err := url.ParseQuery(encodedQuery)
	if err != nil {
		return err
	}

	value, exists := values["pageNumber"]
	if exists && len(value) == 1 {
		page.PageNumber, err = strconv.ParseInt(value[0], 10, 64)
		if err != nil {
			return err
		}
	}

	value, exists = values["pageSize"]
	if exists && len(value) == 1 {
		page.PageSize, err = strconv.ParseInt(value[0], 10, 64)
		if err != nil {
			return err
		}
	}

	return page.Validate()
}

// Check data integrity
func (page PageRequest) Validate() error {

	maxPageSize := settings.GetApplicationProperties().GetRequestProperties().GetMaxPaginationSize()
	if page.PageSize > maxPageSize || page.PageSize < 0 {
		return fmt.Errorf("max page size is %d, min is 0", maxPageSize)
	}

	if page.PageNumber < 0 {
		return errors.New("min page number is 0")
	}
	return nil
}
