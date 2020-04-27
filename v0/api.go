package wdk

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/service"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/path"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/record"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/strategy"
)

type apiProps struct {
	oneSession bool
	sessionId  string
	authToken  string
}

// New instantiates a new WDK API wrapper using the given
// WDK site URL.
//
// See `wdk.NewApiUrl` for the possible errors that may be
// returned
func New(url string) (Api, error) {
	if tmp, err := path.NewApiUrl(url); err != nil {
		return nil, err
	} else {
		return &api{
			apiProps: apiProps{oneSession: true},
			url:      tmp,
			path:     path.NewBuilder(tmp),
		}, nil
	}
}

// ForceNew instantiates a new WDK API wrapper using the
// given WDK site url or panics if any error is encountered.
func ForceNew(url string) Api {
	if out, err := New(url); err != nil {
		panic(err)
	} else {
		return out
	}
}

type Api interface {
	// EnableSessionSharing sets whether or not individual
	// requests should use the same session.
	//
	// Defaults to enabled.
	//
	// This can be disabled and re-enabled for single requests
	// if desired.  After re-enabling, followup requests wil
	// use the session id from before this feature was
	// disabled.
	EnableSessionSharing(bool) Api

	// UseAuthToken sets the QA auth token to use for requests
	// against a QA site.
	UseAuthToken(token string) Api

	// GetUrl returns the resolved URL in use by this API
	// wrapper.
	GetUrl() *path.ApiUrl

	GetServiceDetails() (service.Details, error)

	MustGetServiceDetails() service.Details

	GetPublicStrategyList() (strategy.List, error)

	MustGetPublicStrategyList() strategy.List

	GetRecordTypeNames() (res record.NameList, err error)

	MustGetRecordTypeNames() record.NameList

	GetExpandedRecordTypes() (res record.ExpandedList, err error)

	MustGetExpandedRecordTypes() record.ExpandedList

	UserApiFor(userId uint) UserApi

	CurrentUserApi() UserApi

	RecordApiFor(recordType string) RecordApi
}

type api struct {
	apiProps

	gotSessionId bool

	url  *path.ApiUrl
	path path.Builder
}

func (a *api) ctxLog() *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"shareSessions": a.oneSession,
		"sessionId":     a.sessionId,
		"authToken":     a.authToken,
	})
}

func (a *api) EnableSessionSharing(val bool) Api {
	a.ctxLog().WithField("newSessionShareVal", val).
		Trace("Api.EnableSessionSharing")

	a.oneSession = val

	if !a.gotSessionId && val {
		props := a.apiProps
		props.oneSession = false
		a.sessionId = getSessionId(a.url.String(), a.apiProps)
	}
	return a
}

func (a *api) UseAuthToken(tkn string) Api {
	a.ctxLog().WithField("newTokenVal", tkn).Trace("Api.UseAuthToken")

	a.authToken = tkn

	// Update to a session that's authenticated
	a.sessionId = getSessionId(a.url.String(), a.apiProps)

	return a
}

func (a *api) GetUrl() *path.ApiUrl {
	tmp := *a.url
	return &tmp
}

func (a *api) GetServiceDetails() (res service.Details, err error) {
	a.ctxLog().Trace("Api.GetServiceDetails")

	err = subAndParse(prepGet(a.path.Service(), &a.apiProps), &res)

	if err != nil {
		a.ctxLog().WithField("error", err).
			Debug("Api.GetServiceDetails request failed")
	}

	return
}

func (a *api) MustGetServiceDetails() (res service.Details) {
	a.ctxLog().Trace("Api.MustGetServiceDetails")

	out, err := a.GetServiceDetails()
	if err != nil {
		logger.Panic(err)
	}

	return out
}

func (a *api) GetPublicStrategyList() (res strategy.List, err error) {
	a.ctxLog().Trace("Api.GetPublicStrategyList")

	err = subAndParse(prepGet(a.path.PublicStrategyList(), &a.apiProps), &res)

	if err != nil {
		a.ctxLog().WithField("error", err).
			Debug("Api.GetPublicStrategyList request failed")
	}

	return
}

func (a *api) MustGetPublicStrategyList() strategy.List {
	a.ctxLog().Trace("Api.MustGetPublicStrategyList")

	out, err := a.GetPublicStrategyList()

	if err != nil {
		logger.Panic(err)
	}

	return out
}

func (a *api) GetRecordTypeNames() (res record.NameList, err error) {
	a.ctxLog().Trace("Api.GetRecordTypeNames")

	err = subAndParse(prepGet(a.path.RecordTypes(), &a.apiProps), &res)

	if err != nil {
		a.ctxLog().WithField("error", err).
			Debug("Api.GetRecordTypeNames request failed")
	}

	return
}

func (a *api) MustGetRecordTypeNames() record.NameList {
	a.ctxLog().Trace("Api.MustGetRecordTypeNames")

	out, err := a.GetRecordTypeNames()

	if err != nil {
		logger.Panic(err)
	}

	return out
}

func (a *api) GetExpandedRecordTypes() (res record.ExpandedList, err error) {
	a.ctxLog().Trace("Api.GetExpandedRecordTypes")

	err = subAndParse(prepGet(a.path.RecordTypesExpanded(), &a.apiProps), &res)

	if err != nil {
		a.ctxLog().WithField("error", err).
			Debug("Api.GetExpandedRecordTypes request failed")
	}

	return
}

func (a *api) MustGetExpandedRecordTypes() record.ExpandedList {
	a.ctxLog().Trace("Api.MustGetExpandedRecordTypes")

	out, err := a.GetExpandedRecordTypes()

	if err != nil {
		logger.Panic(err)
	}

	return out
}

func (a *api) UserApiFor(userId uint) UserApi {
	return &userApi{
		builder: path.NewUserBuilder(a.url, userId),
		userId:  strconv.FormatUint(uint64(userId), 10),
		props:   &a.apiProps,
	}
}

func (a *api) CurrentUserApi() UserApi {
	return &userApi{
		builder: path.CurrentPathBuilder(a.url),
		userId:  userCurrent,
		props:   &a.apiProps,
	}
}

func (a *api) RecordApiFor(recordType string) RecordApi {
	return &recordApi{
		recordType: recordType,
		url:        path.NewRecordBuilder(recordType, a.url),
		props:      &a.apiProps,
	}
}
