package path

import (
	"fmt"
	"strconv"
)

const (
	queryExpanded = "format=expanded"

	urlService      = "service"
	urlUsers        = urlService + "/users"
	urlUser         = urlUsers + "/%s"
	urlPublicStrats = urlService + "/strategy-lists/public"
	urlRecordTypes  = urlService + "/record-types"
	urlRecordType   = urlRecordTypes + "/%s"
)

func NewBuilder(url ApiUrl) Builder {
	return &builder{url}
}

type Builder interface {

	// Service returns the URL to the WDK service details
	// endpoint.
	Service() string

	// Users returns the URL to the WDK user list endpoint.
	Users() string

	// User returns the URL to the WDK endpoint providing
	// profile details about the user with the given userId.
	User(userId uint) string

	// CurrentUser returns the URL to the WDK endpoint
	// providing profile details about the currently logged
	// in user.
	CurrentUser() string

	// PublicStrategyList returns the URL to the WDK public
	// strategy list endpoint.
	PublicStrategyList() string

	// RecordTypeNames returns the URL to the WDK endpoint
	// listing the available record types as a list of URL
	// segment strings.
	RecordTypeNames() string

	// RecordTypes returns the URL to the WDK endpoint
	// providing a full listing of available record types and
	// their details.
	RecordTypes() string
}

type builder struct {
	url ApiUrl
}

func (p *builder) Service() string {
	return p.url.wrap(urlService)
}

func (p *builder) Users() string {
	return p.url.wrap(urlUsers)
}

func (p *builder) User(id uint) string {
	return p.url.wrap(fmt.Sprintf(urlUser, strconv.FormatUint(uint64(id), 10)))
}

func (p *builder) CurrentUser() string {
	return p.url.wrap(fmt.Sprintf(urlUser, "current"))
}

func (p *builder) PublicStrategyList() string {
	return p.url.wrap(urlPublicStrats)
}

func (p *builder) RecordTypeNames() string {
	return p.url.wrap(urlRecordTypes)
}

func (p *builder) RecordTypes() string {
	return p.url.wrapQuery(urlRecordTypes, queryExpanded)
}

func (p *builder) RecordType(name string) string {
	return p.url.wrap(fmt.Sprintf(urlRecordType, name))
}
