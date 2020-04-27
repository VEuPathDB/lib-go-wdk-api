package wdk

import (
	"github.com/sirupsen/logrus"

	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/search"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/path"
)

const (
	mnRecordPrefix    = "RecordApi"
	mnGetSearches     = mnRecordPrefix + ".GetSearches"
	mnMustGetSearches = mnRecordPrefix + ".MustGetSearches"
	mnGetSearch       = mnRecordPrefix + ".GetSearch"
	mnMustGetSearch   = mnRecordPrefix + ".MustGetSearch"
)

type RecordApi interface {
	GetSearches() (search.List, error)
	MustGetSearches() search.List

	GetSearch(urlSegment string) (search.FullSearch, error)
	MustGetSearch(urlSegment string) search.FullSearch
}

type recordApi struct {
	recordType string
	url        path.RecordBuilder
	props      *apiProps
}

func (r *recordApi) ctxLog() *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"shareSessions": r.props.oneSession,
		"sessionId":     r.props.sessionId,
		"authToken":     r.props.authToken,
		"recordType":    r.recordType,
	})
}

func (r *recordApi) GetSearches() (res search.List, err error) {
	r.ctxLog().Trace(mnGetSearches)

	err = subAndParse(prepGet(r.url.Searches(), r.props), &res)

	if err != nil {
		r.ctxLog().WithField("error", err).
			Debug(mnGetSearches, "request failed")
	}

	return
}

func (r *recordApi) MustGetSearches() search.List {
	r.ctxLog().Trace(mnMustGetSearches)

	out, err := r.GetSearches()
	if err != nil {
		r.ctxLog().Panic(err)
	}

	return out
}

func (r *recordApi) GetSearch(seg string) (res search.FullSearch, err error) {
	r.ctxLog().Trace(mnGetSearch)

	err = subAndParse(prepGet(r.url.Search(seg), r.props), &res)

	if err != nil {
		r.ctxLog().WithField("error", err).
			Debug(mnGetSearch, "request failed")
	}

	return
}

func (r *recordApi) MustGetSearch(seg string) search.FullSearch {
	r.ctxLog().Trace(mnMustGetSearch)

	out, err := r.GetSearch(seg)
	if err != nil {
		r.ctxLog().Panic(err)
	}

	return out
}
