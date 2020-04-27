package path

import "fmt"

const (
	queryExpanded = "format=expanded"

	urlService      = "service"
	urlUsers        = urlService + "/users"
	urlUser         = urlUsers + "/%s"
	urlPublicStrats = urlService + "/strategy-lists/public"
	urlRecordTypes  = urlService + "/record-types"
	urlRecordType   = urlRecordTypes + "/%s"
)

func NewBuilder(url *ApiUrl) Builder {
	return &builder{url}
}

type Builder interface {

	// Service returns the URL to the WDK service details
	// endpoint.
	Service() string

	// Users returns the URL to the WDK user list endpoint.
	Users() string

	// PublicStrategyList returns the URL to the WDK public
	// strategy list endpoint.
	PublicStrategyList() string

	RecordTypes() string

	RecordTypesExpanded() string
}

type builder struct {
	url *ApiUrl
}

func (p *builder) Service() string {
	return p.url.wrap(urlService)
}

func (p *builder) Users() string {
	return p.url.wrap(urlUsers)
}

func (p *builder) PublicStrategyList() string {
	return p.url.wrap(urlPublicStrats)
}

func (p *builder) RecordTypes() string {
	return p.url.wrap(urlRecordTypes)
}

func (p *builder) RecordTypesExpanded() string {
	return p.url.wrapQuery(urlRecordTypes, queryExpanded)
}

func (p *builder) RecordType(name string) string {
	return p.url.wrap(fmt.Sprintf(urlRecordType, name))
}
