package path

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

func NewBuilder(url ApiUrl) Builder
