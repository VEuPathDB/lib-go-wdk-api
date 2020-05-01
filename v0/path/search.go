package path

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

const (
	errNpbNilApiUrl = "nil ApiUrl instance passed to " + funcNewSearchPathBuilder
)

const (
	urlColumns = "/columns"
	urlColumn  = urlColumns + "/%s"
)

// SearchPathBuilder provides methods for building URLs to
// the search type sub-endpoints
type SearchPathBuilder interface {
	// Columns constructs a URL to the WDK endpoint returning
	// a list of column names for a search.
	Columns() string

	// ColumnsExpanded constructs a URL to the WDK endpoint
	// returning a list of column details for a search.
	ColumnsExpanded() string

	// Column constructs a URL to the WDK endpoint returning
	// details about the specified column.
	Column(columnName string) string

	Url() ApiUrl
}

const funcNewSearchPathBuilder = "NewSearchPathBuilder"

func NewSearchPathBuilder(url ApiUrl, search string) SearchPathBuilder {
	log.Trace(funcNewSearchPathBuilder)
	if url == nil {
		log.Panic(errNpbNilApiUrl)
		return nil
	}

	return &searchBuilder{
		search: search,
		url:    newApiUrl(url.BaseUrl()+fmt.Sprintf(urlSearch, search), url.Query()),
	}
}

type searchBuilder struct {
	search string
	url    ApiUrl
}

func (s *searchBuilder) Url() ApiUrl {
	return s.url
}

func (s *searchBuilder) Columns() string {
	return s.url.wrap(urlColumns)
}

func (s *searchBuilder) ColumnsExpanded() string {
	return s.url.wrapQuery(urlColumns, queryExpanded)
}

func (s *searchBuilder) Column(columnName string) string {
	return s.url.wrap(fmt.Sprintf(urlColumn, columnName))
}
