package wdk

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/service"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/user"
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

const apiClassName = "Api"

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
	GetUrl() path.ApiUrl

	// UrlBuilder returns the internal URL builder used by
	// this API class.
	UrlBuilder() path.Builder

	// GetServiceDetails retrieves and unmarshals the response
	// from the WDK API details endpoint and returns either
	// the result on success or an error if the request or
	// unmarshaling failed.
	GetServiceDetails() (service.Details, error)

	// MustGetServiceDetails performs the same steps as
	// GetServiceDetails except it panics on error.
	MustGetServiceDetails() service.Details

	GetPublicStrategyList() (strategy.List, error)
	MustGetPublicStrategyList() strategy.List

	GetRecordTypeNames() ([]string, error)
	MustGetRecordTypeNames() []string

	GetRecordTypes() ([]record.Type, error)
	MustGetRecordTypes() []record.Type

	GetUser(userId uint) (user.Profile, error)
	MustGetUser(userId uint) user.Profile
	GetCurrentUser() (user.Profile, error)
	MustGetCurrentUser() user.Profile

	// UserApiFor returns an API access type for dealing with
	// API endpoints dependent on a specific user.
	UserApiFor(userId uint) UserApi
	CurrentUserApi() UserApi

	RecordApiFor(recordType string) RecordApi
}

type api struct {
	apiProps

	gotSessionId bool

	url  path.ApiUrl
	path path.Builder
}

const apiReqFailed = " request failed"

func (a *api) ctxLog() *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"shareSessions": a.oneSession,
		"sessionId":     a.sessionId,
		"authToken":     a.authToken,
	})
}

func (a *api) UrlBuilder() path.Builder {
	return a.path
}

func (a *api) EnableSessionSharing(val bool) Api {
	a.ctxLog().WithField("newSessionShareVal", val).
		Trace("Api.EnableSessionSharing")

	a.oneSession = val

	if !a.gotSessionId && val {
		props := a.apiProps
		props.oneSession = false
		a.sessionId = getSessionId(a.path.Service(), a.apiProps)
	}
	return a
}

func (a *api) UseAuthToken(tkn string) Api {
	a.ctxLog().WithField("newTokenVal", tkn).Trace("Api.UseAuthToken")

	a.authToken = tkn

	// Update to a session that's authenticated
	a.sessionId = getSessionId(a.path.Service(), a.apiProps)

	return a
}

func (a *api) GetUrl() path.ApiUrl {
	return a.url
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

func (a *api) GetRecordTypeNames() (res []string, err error) {
	a.ctxLog().Trace("Api.GetRecordTypeNames")

	err = subAndParse(prepGet(a.path.RecordTypeNames(), &a.apiProps), &res)

	if err != nil {
		a.ctxLog().WithField("error", err).
			Debug("Api.GetRecordTypeNames request failed")
	}

	return
}

func (a *api) MustGetRecordTypeNames() []string {
	a.ctxLog().Trace("Api.MustGetRecordTypeNames")

	out, err := a.GetRecordTypeNames()

	if err != nil {
		logger.Panic(err)
	}

	return out
}

const apiMethGetRecTypes = apiClassName + ".GetRecordTypes"

func (a *api) GetRecordTypes() (res []record.Type, err error) {
	a.ctxLog().Trace(apiMethGetRecTypes)

	err = subAndParse(prepGet(a.path.RecordTypes(), &a.apiProps), &res)

	if err != nil {
		a.ctxLog().
			WithField("error", err).
			Debug(apiMethGetRecTypes + apiReqFailed)
	}

	return
}

const apiMethMustGetRecTypes = apiClassName + ".MustGetRecordTypes"

func (a *api) MustGetRecordTypes() []record.Type {
	a.ctxLog().Trace(apiMethMustGetRecTypes)

	out, err := a.GetRecordTypes()

	if err != nil {
		a.ctxLog().Panic(err)
	}

	return out
}

const apiMethGetUser = apiClassName + ".GetUser"

func (a *api) GetUser(userId uint) (res user.Profile, err error) {
	a.ctxLog().Trace(apiMethGetUser)

	err = subAndParse(prepGet(a.path.User(userId), &a.apiProps), &res)

	if err != nil {
		a.ctxLog().WithField("error", err).Debug(apiMethGetUser + apiReqFailed)
	}

	return
}

const apiMethMustGetUser = apiClassName + ".MustGetUser"

func (a *api) MustGetUser(userId uint) user.Profile {
	a.ctxLog().Trace(apiMethMustGetUser)

	out, err := a.GetUser(userId)

	if err != nil {
		a.ctxLog().Panic(err)
	}

	return out
}

const apiMethGetCurUser = apiClassName + ".GetCurrentUser"

func (a *api) GetCurrentUser() (res user.Profile, err error) {
	a.ctxLog().Trace(apiMethGetCurUser)

	err = subAndParse(prepGet(a.path.CurrentUser(), &a.apiProps), &res)

	if err != nil {
		a.ctxLog().WithField("error", err).Debug(apiMethGetUser + apiReqFailed)
	}

	return
}

const apiMethMustGetCurUser = apiClassName + ".MustGetCurrentUser"

func (a *api) MustGetCurrentUser() user.Profile {
	a.ctxLog().Trace(apiMethMustGetCurUser)

	out, err := a.GetCurrentUser()

	if err != nil {
		a.ctxLog().Panic(err)
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
		rType: recordType,
		path:  path.NewRecordBuilder(a.url, recordType),
		props: &a.apiProps,
	}
}
