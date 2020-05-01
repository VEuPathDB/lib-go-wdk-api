package path

import "fmt"

const (
	urlSearches = "/searches"
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

	Url() ApiUrl
}

// NewRecordBuilder constructs a new path builder
// specifically for record type sub-paths.
func NewRecordBuilder(base ApiUrl, recordType string) RecordPathBuilder {
	return &recordBuilder{
		recordType: recordType,
		url:        newApiUrl(base.BaseUrl()+fmt.Sprintf(urlRecordType, recordType), base.Query()),
	}
}

type recordBuilder struct {
	recordType string
	url        ApiUrl
}

func (r *recordBuilder) Url() ApiUrl {
	return r.url
}

func (r *recordBuilder) Searches() string {
	return r.url.wrap(urlSearches)
}

func (r *recordBuilder) Search(search string) string {
	return r.url.wrap(fmt.Sprintf(urlSearch, search))
}
