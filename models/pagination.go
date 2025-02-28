package models

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/lucas-10101/logapi/web/webutils"
)

type Pagination struct {
	PageNumber int
	PageSize   int
}

// parse url variables 'pagenumber' and 'pagesize' into *Pagination struct
func GetPaginationFromUrl(urlData *url.URL) (*Pagination, webutils.HttpError) {

	values, err := url.ParseQuery(urlData.RawQuery)
	if err != nil {
		return nil, webutils.NewHttpError(http.StatusBadRequest, "url pagination data is invalid")
	}
	paginationData := &Pagination{}

	queryValue, queryExists := values["pagenumber"]
	if queryExists && len(queryValue) == 1 {
		if paginationData.PageNumber, err = strconv.Atoi(queryValue[0]); err != nil || paginationData.PageNumber < 0 {
			return nil, webutils.NewHttpError(http.StatusBadRequest, "url pagination 'pagenumber' is invalid")
		}
	}

	queryValue, queryExists = values["pagesize"]
	if queryExists && len(queryValue) == 1 {
		if paginationData.PageSize, err = strconv.Atoi(queryValue[0]); err != nil || paginationData.PageSize > 50 {
			return nil, webutils.NewHttpError(http.StatusBadRequest, "url pagination 'pagesize' is invalid")
		}
	}
	return paginationData, nil
}
