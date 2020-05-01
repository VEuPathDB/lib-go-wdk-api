package wdk

import (
	"github.com/sirupsen/logrus"

	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/search"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/path"
)

// type/method names for logging
const (
	mnRecordPrefix    = "RecordApi"
	mnGetSearches     = mnRecordPrefix + ".GetSearches"
	mnMustGetSearches = mnRecordPrefix + ".MustGetSearches"
	mnGetSearch       = mnRecordPrefix + ".GetSearch"
	mnMustGetSearch   = mnRecordPrefix + ".MustGetSearch"
)

// RecordApi contains methods for dealing with the
// record-type WDK sub-api.
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

	// PostRecords() (record._____, error)
	// MustPostRecords() record._____

	// SearchApiFor takes the UrlSegment for a search and
	// constructs an API wrapper for the searches API
	// sub-endpoints.
	SearchApiFor(searchName string) SearchApi
}

type recordApi struct {
	rType string
	path  path.RecordPathBuilder
	props *apiProps
}

func (r *recordApi) ctxLog() *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"shareSessions": r.props.oneSession,
		"sessionId":     r.props.sessionId,
		"authToken":     r.props.authToken,
		"recordType":    r.rType,
	})
}

func (r *recordApi) UrlBuilder() path.RecordPathBuilder {
	return r.path
}

func (r *recordApi) GetSearches() (res []search.ShortSearch, err error) {
	r.ctxLog().Trace(mnGetSearches)

	err = subAndParse(prepGet(r.path.Searches(), r.props), &res)

	if err != nil {
		r.ctxLog().WithField("error", err).
			Debug(mnGetSearches, "request failed")
	}

	return
}

func (r *recordApi) MustGetSearches() []search.ShortSearch {
	r.ctxLog().Trace(mnMustGetSearches)

	out, err := r.GetSearches()
	if err != nil {
		r.ctxLog().Panic(err)
	}

	return out
}

func (r *recordApi) GetSearch(seg string) (res search.ValidatedSearch, err error) {
	r.ctxLog().Trace(mnGetSearch)

	err = subAndParse(prepGet(r.path.Search(seg), r.props), &res)

	if err != nil {
		r.ctxLog().WithField("error", err).
			Debug(mnGetSearch, "request failed")
	}

	return
}

func (r *recordApi) MustGetSearch(seg string) search.ValidatedSearch {
	r.ctxLog().Trace(mnMustGetSearch)

	out, err := r.GetSearch(seg)
	if err != nil {
		r.ctxLog().Panic(err)
	}

	return out
}

func (r *recordApi) SearchApiFor(name string) SearchApi {
	return &searchApi{
		record: r.rType,
		search: name,
		path:   path.NewSearchPathBuilder(r.path.Url(), name),
		props:  nil,
	}
}
