package wdk

import (
	"github.com/veupathdb/lib-go-rest-types/pkg/veupath/path"

	"lib-go-wdk-api/pkg/wdk/model/record"
	"lib-go-wdk-api/pkg/wdk/model/service"
	"lib-go-wdk-api/pkg/wdk/model/strategy"
	"lib-go-wdk-api/pkg/wdk/model/user"
)

type Api interface {
	// EnableSessionSharing sets whether individual requests should use the same
	// session.
	//
	// Defaults to enabled.
	//
	// This can be disabled and re-enabled for single requests if desired.  After
	// re-enabling, followup requests wil use the session id from before this
	// feature was disabled.
	EnableSessionSharing(bool) Api

	// UseAuthToken sets the QA auth token to use for requests against a QA site.
	UseAuthToken(token string) Api

	// GetUrl returns the resolved URL in use by this API wrapper.
	GetUrl() path.ApiUrl

	// UrlBuilder returns the internal URL builder used by this API class.
	UrlBuilder() path.Builder

	// GetServiceDetails retrieves and unmarshals the response from the WDK API
	// details endpoint and returns either the result on success or an error if
	// the request or unmarshaling failed.
	GetServiceDetails() (service.Details, error)

	// MustGetServiceDetails performs the same steps as GetServiceDetails except
	// it panics on error.
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

	// UserApiFor returns an API access type for dealing with API endpoints
	// dependent on a specific user.
	UserApiFor(userId uint) UserApi
	CurrentUserApi() UserApi

	RecordApiFor(recordType string) RecordApi
}

func ForceNew(url string) Api {

}

func New(url string) (Api, error) {

}
