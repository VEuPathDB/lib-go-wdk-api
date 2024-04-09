package path

import "fmt"

// ApiUrl wraps a validated URL base and query for internal use.
type ApiUrl interface {
	fmt.Stringer

	BaseUrl() string
	Query() string
	// contains filtered or unexported methods
}

// NewApiUrl parses the input url and returns an internal representation of that URL data.
//
// An error is returned if the url given does not appear to be valid or if the site it points to is not reachable.
//
// Valid input url formats:
//
//   site.com
//   site.com?some=query&string=value
//   site.com/app
//   site.com/app?some=query&string=value
//   http://site.com
//   http://site.com?some=query&string=value
//   http://site.com/app
//   http://site.com/app?some=query&string=value
//   https://site.com
//   https://site.com?some=query&string=value
//   https://site.com/app
//   https://site.com/app?some=query&string=value
func NewApiUrl(url string) (ApiUrl, error)
