package wdk

const (
	urlService      = "service"
	urlUsers        = urlService + "/users"
	urlUser         = urlUsers + "/%s"
	urlPublicStrats = urlService + "/strategy-lists/public"
)

func NewPathBuilder(url *ApiUrl) PathBuilder {
	return &pathBuilder{url}
}

type PathBuilder interface {

	// Service returns the URL to the WDK service details
	// endpoint.
	Service() string

	// Users returns the URL to the WDK user list endpoint.
	Users() string

	// PublicStrategyList returns the URL to the WDK public
	// strategy list endpoint.
	PublicStrategyList() string
}

type pathBuilder struct {
	url *ApiUrl
}

func (p *pathBuilder) Service() string {
	return p.url.wrap(urlService)
}

func (p *pathBuilder) Users() string {
	return p.url.wrap(urlUsers)
}

func (p *pathBuilder) PublicStrategyList() string {
	return p.url.wrap(urlPublicStrats)
}
