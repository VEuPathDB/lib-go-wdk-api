package wdk

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Foxcapades/Go-ChainRequest/simple"
)

const (
	locationHeader = "Location"
)

// NewApiUrl parses the input url and returns an internal
// representation of that URL data.
//
// An error is returned if the url given does not appear to
// be valid or if the site it points to is not reachable.
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
func NewApiUrl(url string) (*ApiUrl, error) {
	out := new(ApiUrl)

	if err := out.parseUrl(url); err != nil {
		return nil, err
	}

	if err := out.resolve(); err != nil {
		return nil, err
	}

	return out, nil
}

type ApiUrl struct {
	base  string
	query string
	hasQ  bool
}

func (a *ApiUrl) String() string {
	return a.base + a.query
}

func (a *ApiUrl) wrap(in string) string {
	if a.hasQ {
		return a.base + in + a.query
	} else {
		return a.base + in
	}
}

func (a *ApiUrl) wrapQuery(in, q string) string {
	if a.hasQ {
		return a.base + in + a.query + "&" + q
	} else {
		return a.base + in + "?" + q
	}
}

func (a *ApiUrl) parseUrl(url string) error {
	if i := strings.IndexByte(url, '?'); i > -1 {
		a.query = url[i:]
		a.hasQ  = true
		url = url[:i]
	}

	if len(url) < 4 {
		return errors.New("url too short to be valid")
	}

	if !strings.HasSuffix(url, "/") {
		url += "/"
	}

	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	} else if url[4] != 's' {
		url = strings.Replace(url, "http:", "https:", 1)
	}
	a.base = url

	return nil
}

func (a *ApiUrl) resolve() error {
	res := simple.GetRequest(a.String()).DisableRedirects().Submit()
	defer res.Close()

	if err := res.GetError(); err != nil {
		return err
	}

	if res.MustGetResponseCode() == http.StatusFound {
		location := res.MustGetHeader(locationHeader)

		if strings.Index(location, "/app/") == -1 {
			return a.parseUrl(location)
		}
	}

	return nil
}
