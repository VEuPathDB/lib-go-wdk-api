package wdk

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/read/users/strategies"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/write"
)

type UserApi interface {
	// GetStrategies returns a listing of strategies owned by
	// the current user.
	GetStrategies() strategies.ListResponse

	// CopyStrategy submits a strategy copy request using the
	// given signature as the source strategy reference.
	//
	// The current user must have access to a strategy
	// matching that signature or this will fail.
	CopyStrategy(signature string) (strategies.CopyResponse, error)

	// GetStrategy looks up the strategy with the given id.
	GetStrategy(strategyId uint) (strategies.Strategy, error)
}

type userApi struct {
	builder UserPathBuilder
	userId  string
	props   *apiProps
}

func (u *userApi) GetStrategies() strategies.ListResponse {
	var res strategies.ListResponse
	prepGet(u.builder.Strategies(), u.props).Submit().
		MustUnmarshalBody(res, unmarshaler)
	return res
}

func (u *userApi) CopyStrategy(signature string) (res strategies.CopyResponse, err error) {
	body := write.CopyStrategyReq{SourceStrategySignature: signature}
	err = prepPost(u.builder.Strategies(), u.props).
		MarshalBody(body, marshaler).
		Submit().
		UnmarshalBody(&res, unmarshaler)
	return
}

func (u *userApi) GetStrategy(id uint) (res strategies.Strategy, err error) {
	err = prepGet(u.builder.Strategy(id), u.props).Submit().
		UnmarshalBody(&res, unmarshaler)
	return
}
