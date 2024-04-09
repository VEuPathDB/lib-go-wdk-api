package wdk

type RecordApi interface {
	// UrlBuilder returns the internal URL builder used by
	// this API class.
	UrlBuilder() path.RecordPathBuilder

	// GetSearches returns the list of available searches for
	// the current record type, or an error if the request
	// failed or response could not be parsed.
	GetSearches() ([]search.ShortSearch, error)

	// MustGetSearches does the same as GetSearches except
	// it will panic on error.
	MustGetSearches() []search.ShortSearch

	// GetSearch details about the given search from the
	// current record type, or an error if the request failed
	// or the response could not be parsed.
	GetSearch(urlSegment string) (search.ValidatedSearch, error)

	// MustGetSearch does the same as GetSearch except it will
	// panic on error.
	MustGetSearch(urlSegment string) search.ValidatedSearch

	// SearchApiFor takes the UrlSegment for a search and
	// constructs an API wrapper for the searches API
	// sub-endpoints.
	SearchApiFor(searchName string) SearchApi
}
