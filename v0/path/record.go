package path

import "fmt"

const (
	urlSearches = urlRecordTypes + "/%s/searches"
	urlSearch   = urlSearches + "/%s"
)

type RecordBuilder interface {
	Searches() string
	Search(search string) string
}

type recordBuilder struct {
	recordType string
	url        *ApiUrl
}

func NewRecordBuilder(recordType string, url *ApiUrl) RecordBuilder {
	return &recordBuilder{
		recordType: recordType,
		url:        url,
	}
}

func (r *recordBuilder) Searches() string {
	return r.url.wrap(fmt.Sprintf(urlSearches, r.recordType))
}

func (r *recordBuilder) Search(search string) string {
	return r.url.wrap(fmt.Sprintf(urlSearch, r.recordType, search))
}
