package path

// RecordPathBuilder provides methods for building URLs to the record type sub-endpoints.
type RecordPathBuilder interface {

	// Searches constructs a URL to reach the search list
	// REST endpoint.
	Searches() string

	// Search constructs a URL to reach the search lookup by
	// name REST endpoint using the given search name.
	Search(search string) string

	Url() ApiUrl
}

// NewRecordPathBuilder constructs a new path builder specifically for record type sub-paths.
func NewRecordPathBuilder(base ApiUrl, recordType string) RecordPathBuilder
