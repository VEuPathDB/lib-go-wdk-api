package wdk

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/recordTypes"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/VEuPathDB/lib-go-wdk-api/v0/service"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/strategy-lists/public"
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
	if tmp, err := NewApiUrl(url); err != nil {
		return nil, err
	} else {
		return &api{
			apiProps: apiProps{oneSession: true},
			url:      tmp,
			path:     NewPathBuilder(tmp),
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
	GetUrl() *ApiUrl

	GetServiceDetails() (service.Service, error)

	MustGetServiceDetails() service.Service

	GetPublicStrategyList() (public.StrategyList, error)

	MustGetPublicStrategyList() public.StrategyList

	GetRecordTypeNames() (res recordTypes.NameList, err error)

	MustGetRecordTypeNames() recordTypes.NameList

	GetExpandedRecordTypes() (res recordTypes.ExpandedList, err error)

	MustGetExpandedRecordTypes() recordTypes.ExpandedList

	UserApiFor(userId uint) UserApi

	CurrentUserApi() UserApi
}

type api struct {
	apiProps

	gotSessionId bool

	url  *ApiUrl
	path PathBuilder
}

func (a *api) EnableSessionSharing(val bool) Api {
	logger.WithField("sessionSharing", val).Trace("Api.EnableSessionSharing")

	a.oneSession = val

	if !a.gotSessionId && val {
		props := a.apiProps
		props.oneSession = false
		a.sessionId = getSessionId(a.url.String(), a.apiProps)
	}
	return a
}

func (a *api) UseAuthToken(tkn string) Api {
	logger.WithField("token", tkn).Trace("Api.UseAuthToken")

	a.authToken = tkn

	// Update to a session that's authenticated
	a.sessionId = getSessionId(a.url.String(), a.apiProps)

	return a
}

func (a *api) GetUrl() *ApiUrl {
	tmp := *a.url
	return &tmp
}

func (a *api) GetServiceDetails() (res service.Service, err error) {
	ctxLog := logger.WithFields(logrus.Fields{
		"shareSessions": a.oneSession,
		"sessionId":     a.sessionId,
		"authToken":     a.authToken,
	})
	ctxLog.Trace("Api.GetServiceDetails")

	err = subAndParse(prepGet(a.path.Service(), &a.apiProps), &res)

	if err != nil {
		ctxLog.WithField("error", err).Debug("Api.GetServiceDetails request failed")
	}

	return
}

func (a *api) MustGetServiceDetails() (res service.Service) {
	logger.WithFields(logrus.Fields{
		"shareSessions": a.oneSession,
		"sessionId":     a.sessionId,
		"authToken":     a.authToken,
	}).Trace("Api.MustGetServiceDetails")

	out, err := a.GetServiceDetails()
	if err != nil {
		logger.Panic(err)
	}

	return out
}

func (a *api) GetPublicStrategyList() (res public.StrategyList, err error) {
	ctxLog := logger.WithFields(logrus.Fields{
		"shareSessions": a.oneSession,
		"sessionId":     a.sessionId,
		"authToken":     a.authToken,
	})
	ctxLog.Trace("Api.GetPublicStrategyList")

	err = subAndParse(prepGet(a.path.PublicStrategyList(), &a.apiProps), &res)

	if err != nil {
		ctxLog.WithField("error", err).Debug("Api.GetPublicStrategyList request failed")
	}

	return
}

func (a *api) MustGetPublicStrategyList() public.StrategyList {
	logger.WithFields(logrus.Fields{
		"shareSessions": a.oneSession,
		"sessionId":     a.sessionId,
		"authToken":     a.authToken,
	}).Trace("Api.MustGetPublicStrategyList")

	out, err := a.GetPublicStrategyList()

	if err != nil {
		logger.Panic(err)
	}

	return out
}

func (a *api) GetRecordTypeNames() (res recordTypes.NameList, err error) {
	ctxLog := logger.WithFields(logrus.Fields{
		"shareSessions": a.oneSession,
		"sessionId":     a.sessionId,
		"authToken":     a.authToken,
	})
	ctxLog.Trace("Api.GetRecordTypeNames")

	err = subAndParse(prepGet(a.path.RecordTypes(), &a.apiProps), &res)

	if err != nil {
		ctxLog.WithField("error", err).Debug("Api.GetRecordTypeNames request failed")
	}

	return
}

func (a *api) MustGetRecordTypeNames() recordTypes.NameList {
	logger.WithFields(logrus.Fields{
		"shareSessions": a.oneSession,
		"sessionId":     a.sessionId,
		"authToken":     a.authToken,
	}).Trace("Api.MustGetRecordTypeNames")

	out, err := a.GetRecordTypeNames()

	if err != nil {
		logger.Panic(err)
	}

	return out
}

func (a *api) GetExpandedRecordTypes() (res recordTypes.ExpandedList, err error) {
	ctxLog := logger.WithFields(logrus.Fields{
		"shareSessions": a.oneSession,
		"sessionId":     a.sessionId,
		"authToken":     a.authToken,
	})
	ctxLog.Trace("Api.GetExpandedRecordTypes")

	err = subAndParse(prepGet(a.path.RecordTypesExpanded(), &a.apiProps), &res)

	if err != nil {
		ctxLog.WithField("error", err).Debug("Api.GetExpandedRecordTypes request failed")
	}

	return
}

func (a *api) MustGetExpandedRecordTypes() recordTypes.ExpandedList {
	logger.WithFields(logrus.Fields{
		"shareSessions": a.oneSession,
		"sessionId":     a.sessionId,
		"authToken":     a.authToken,
	}).Trace("Api.MustGetExpandedRecordTypes")

	out, err := a.GetExpandedRecordTypes()

	if err != nil {
		logger.Panic(err)
	}

	return out
}

func (a *api) UserApiFor(userId uint) UserApi {
	return &userApi{
		builder: NewUserPathBuilder(a.url, userId),
		userId:  strconv.FormatUint(uint64(userId), 10),
		props:   &a.apiProps,
	}
}

func (a *api) CurrentUserApi() UserApi {
	return &userApi{
		builder: CurrentUserPathBuilder(a.url),
		userId:  userCurrent,
		props:   &a.apiProps,
	}
}
