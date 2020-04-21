package wdk

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/read/users/strategies"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/write"
	log "github.com/sirupsen/logrus"
)

type UserApi interface {
	// GetStrategies returns a listing of strategies owned by
	// the current user.
	GetStrategies() (strategies.ListResponse, error)

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

func (u *userApi) GetStrategies() (res strategies.ListResponse, err error) {
	ctxLog := logger.WithFields(log.Fields{
		"shareSessions": u.props.oneSession,
		"sessionId":     u.props.sessionId,
		"authToken":     u.props.authToken,
	})
	ctxLog.Trace("UserApi.GetStrategies")

	err = prepGet(u.builder.Strategies(), u.props).Submit().
		UnmarshalBody(&res, unmarshaler)

	if err != nil {
		ctxLog.WithField("error", err).Debug("UserApi.GetStrategies request failed")
	}

	return
}

func (u *userApi) CopyStrategy(signature string) (res strategies.CopyResponse, err error) {
	ctxLog := logger.WithFields(log.Fields{
		"shareSessions": u.props.oneSession,
		"sessionId":     u.props.sessionId,
		"authToken":     u.props.authToken,
		"signature":     signature,
	})
	ctxLog.Trace("UserApi.CopyStrategy")

	body := write.CopyStrategyReq{SourceStrategySignature: signature}
	err = prepPost(u.builder.Strategies(), u.props).
		MarshalBody(body, marshaler).
		Submit().
		UnmarshalBody(&res, unmarshaler)

	if err != nil {
		ctxLog.WithField("error", err).Debug("UserApi.CopyStrategy request failed")
	}

	return
}

func (u *userApi) GetStrategy(id uint) (res strategies.Strategy, err error) {
	ctxLog := logger.WithFields(log.Fields{
		"shareSessions": u.props.oneSession,
		"sessionId":     u.props.sessionId,
		"authToken":     u.props.authToken,
		"strategyId":    id,
	})
	ctxLog.Trace("UserApi.GetStrategy")

	err = prepGet(u.builder.Strategy(id), u.props).Submit().
		UnmarshalBody(&res, unmarshaler)

	if err != nil {
		ctxLog.WithField("error", err).Debug("UserApi.GetStrategy request failed")
	}

	return
}
