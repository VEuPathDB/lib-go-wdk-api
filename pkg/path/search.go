package path

// SearchPathBuilder provides methods for building URLs to the search type sub-endpoints
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

func NewSearchPathBuilder(url ApiUrl, search string) SearchPathBuilder
