package wdk

type UserApi interface {
	// UrlBuilder returns the internal URL builder used by
	// this API class.
	UrlBuilder() path.UserPathBuilder

	// GetStrategies returns a listing of strategies owned by
	// the current user.
	GetStrategies() (strategy.List, error)

	MustGetStrategies() strategy.List

	// CopyStrategy submits a strategy copy request using the
	// given signature as the source strategy reference.
	//
	// The current user must have access to a strategy
	// matching that signature or this will fail.
	CopyStrategy(signature string) (strategy.CopyResponse, error)

	MustCopyStrategy(signature string) strategy.CopyResponse

	// GetStrategy looks up the strategy with the given id.
	GetStrategy(strategyId uint) (strategy.FullStrategy, error)

	MustGetStrategy(strategyId uint) strategy.FullStrategy
}
