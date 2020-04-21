package wdk

import (
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

	url  *ApiUrl
	path PathBuilder
}

func (a *api) EnableSessionSharing(val bool) Api {
	a.oneSession = val
	a.sessionId = getSessionId(a.url.String())
	return a
}

func (a *api) UseAuthToken(tkn string) Api {
	a.authToken = tkn
	return a
}

func (a *api) GetServiceDetails() (res read.Service, err error) {
	err = prepGet(a.path.Service(), &a.apiProps).Submit().
		UnmarshalBody(&res, unmarshaler)
	return
}

func (a *api) MustGetServiceDetails() (res read.Service) {
	prepGet(a.path.Service(), &a.apiProps).
		Submit().
		MustUnmarshalBody(&res, unmarshaler)
	return
}

func (a *api) GetPublicStrategyList() (res public.StrategyList, err error) {
	err = prepGet(a.path.PublicStrategyList(), &a.apiProps).Submit().
		UnmarshalBody(&res, unmarshaler)
	return
}

func (a *api) MustGetPublicStrategyList() (res public.StrategyList) {
	prepGet(a.path.PublicStrategyList(), &a.apiProps).
		Submit().
		MustUnmarshalBody(&res, unmarshaler)
	return
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
