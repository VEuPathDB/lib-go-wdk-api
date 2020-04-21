package wdk

import (
	log "github.com/sirupsen/logrus"
	"strconv"

	"github.com/VEuPathDB/lib-go-wdk-api/v0/read"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/read/strategy-lists/public"
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

	GetServiceDetails() (read.Service, error)

	MustGetServiceDetails() read.Service

	GetPublicStrategyList() (public.StrategyList, error)

	MustGetPublicStrategyList() public.StrategyList

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

func (a *api) GetServiceDetails() (res read.Service, err error) {
	ctxLog := logger.WithFields(log.Fields{
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

func (a *api) MustGetServiceDetails() (res read.Service) {
	logger.WithFields(log.Fields{
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
	ctxLog := logger.WithFields(log.Fields{
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
	logger.WithFields(log.Fields{
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
