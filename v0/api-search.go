package wdk

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/attribute"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/path"
	"github.com/sirupsen/logrus"
)

// type/method names for logging
const (
	searchApiClass    = "SearchApi"
	mnGetColNames     = searchApiClass + ".GetColumnNames"
	mnMustGetColNames = searchApiClass + ".MustGetColumnNames"
	mnhGetCols        = searchApiClass + ".GetColumns"
	mnMustGetCols     = searchApiClass + ".MustGetColumns"
)

type SearchApi interface {
	GetColumnNames() ([]string, error)
	MustGetColumnNames() []string

	GetColumns() ([]attribute.Field, error)
	MustGetColumns() []attribute.Field

	// TODO: what should this method be named?
	// PostRevise() (search._____, error)
	// MustPostRevise() search._____

	// PostSearchChange() (search._____, error)
	// MustPostSearchChange() search._____

	// PostFilterCounts() (search._____, error)
	// MustPostFilterCounts() search._____

	ColumnApiFor(columnName string) ColumnApi
}

type searchApi struct {
	record string
	search string
	path   path.SearchPathBuilder
	props  *apiProps
}

func (s *searchApi) ctxLog() *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"shareSessions": s.props.oneSession,
		"sessionId":     s.props.sessionId,
		"authToken":     s.props.authToken,
		"recordType":    s.record,
		"search":        s.search,
	})
}

func (s *searchApi) GetColumnNames() (res []string, err error) {
	s.ctxLog().Trace(mnGetColNames)
	logFail(subAndParse(prepGet(s.path.Columns(), s.props), &res),
		s.ctxLog(), mnGetColNames)
	return
}

func (s *searchApi) MustGetColumnNames() []string {
	s.ctxLog().Trace(mnMustGetColNames)

	out, err := s.GetColumnNames()
	if err != nil {
		s.ctxLog().Panic(err)
	}

	return out
}

func (s *searchApi) GetColumns() (res []attribute.Field, err error) {
	s.ctxLog().Trace(mnhGetCols)
	logFail(subAndParse(prepGet(s.path.Columns(), s.props), &res),
		s.ctxLog(), mnhGetCols)
	return
}

func (s *searchApi) MustGetColumns() []attribute.Field {
	s.ctxLog().Trace(mnMustGetCols)

	out, err := s.GetColumns()
	if err != nil {
		s.ctxLog().Panic(err)
	}

	return out
}

func (s *searchApi) ColumnApiFor(name string) ColumnApi {
	panic("implement me")
}
