package models

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/lucas-10101/logapi/web/webutils"
)

type Pagination struct {
	PageNumber int64
	PageSize   int64
}

// parse url variables 'pagenumber' and 'pagesize' into *Pagination struct
func GetPaginationFromUrl(urlData *url.URL) (*Pagination, webutils.HttpError) {

	values, err := url.ParseQuery(urlData.RawQuery)
	if err != nil {
		return nil, webutils.NewHttpError(http.StatusBadRequest, "url pagination data is invalid")
	}
	paginationData := &Pagination{
		PageNumber: 0,
		PageSize:   30,
	}

	queryValue, queryExists := values["pagenumber"]
	if queryExists && len(queryValue) == 1 {

		paginationData.PageNumber, err = strconv.ParseInt(queryValue[0], 10, 64)
		if err != nil || paginationData.PageNumber < 0 {
			return nil, webutils.NewHttpError(http.StatusBadRequest, "url pagination 'pagenumber' is invalid")
		}
	}

	queryValue, queryExists = values["pagesize"]
	if queryExists && len(queryValue) == 1 {

		paginationData.PageSize, err = strconv.ParseInt(queryValue[0], 10, 64)
		if err != nil || paginationData.PageSize < 1 || paginationData.PageSize > 50 {
			return nil, webutils.NewHttpError(http.StatusBadRequest, "url pagination 'pagesize' is invalid")
		}
	}
	return paginationData, nil
}
