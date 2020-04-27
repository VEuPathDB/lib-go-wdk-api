package path

import "fmt"

const (
	urlSearches = urlRecordTypes + "/%s/searches"
	urlSearch   = urlSearches + "/%s"
)

// RecordPathBuilder provides methods for building URLs to the
// record type sub-endpoints.
type RecordPathBuilder interface {

	// Searches constructs a URL to reach the search list
	// REST endpoint.
	Searches() string

	// Search constructs a URL to reach the search lookup by
	// name REST endpoint using the given search name.
	Search(search string) string
}

// NewRecordBuilder constructs a new path builder
// specifically for record type sub-paths.
func NewRecordBuilder(base ApiUrl, recordType string) RecordPathBuilder {
	return &recordBuilder{
		recordType: recordType,
		url:        newApiUrl(base.BaseUrl() + fmt.Sprintf(urlRecordType, recordType)),
	}
}

type recordBuilder struct {
	recordType string
	url        ApiUrl
}

func (r *recordBuilder) Searches() string {
	return r.url.wrap(fmt.Sprintf(urlSearches, r.recordType))
}

func (r *recordBuilder) Search(search string) string {
	return r.url.wrap(fmt.Sprintf(urlSearch, r.recordType, search))
}
