package wdk

import (
	"github.com/sirupsen/logrus"

	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/users/strategies"
)

type UserApi interface {
	// GetStrategies returns a listing of strategies owned by
	// the current user.
	GetStrategies() (strategies.ListResponse, error)

	MustGetStrategies() strategies.ListResponse

	// CopyStrategy submits a strategy copy request using the
	// given signature as the source strategy reference.
	//
	// The current user must have access to a strategy
	// matching that signature or this will fail.
	CopyStrategy(signature string) (strategies.CopyResponse, error)

	MustCopyStrategy(signature string) strategies.CopyResponse

	// GetStrategy looks up the strategy with the given id.
	GetStrategy(strategyId uint) (strategies.Strategy, error)

	MustGetStrategy(strategyId uint) strategies.Strategy
}

type userApi struct {
	builder UserPathBuilder
	userId  string
	props   *apiProps
}

func (u *userApi) GetStrategies() (res strategies.ListResponse, err error) {
	ctxLog := logger.WithFields(logrus.Fields{
		"shareSessions": u.props.oneSession,
		"sessionId":     u.props.sessionId,
		"authToken":     u.props.authToken,
	})
	ctxLog.Trace("UserApi.GetStrategies")

	err = subAndParse(prepGet(u.builder.Strategies(), u.props), &res)

	if err != nil {
		ctxLog.WithField("error", err).Debug("UserApi.GetStrategies request failed")
	}

	return
}

func (u *userApi) MustGetStrategies() (res strategies.ListResponse) {
	logger.WithFields(logrus.Fields{
		"shareSessions": u.props.oneSession,
		"sessionId":     u.props.sessionId,
		"authToken":     u.props.authToken,
	}).Trace("Api.MustGetStrategies")

	out, err := u.GetStrategies()
	if err != nil {
		logger.Panic(err)
	}

	return out
}


func (u *userApi) CopyStrategy(signature string) (res strategies.CopyResponse, err error) {
	ctxLog := logger.WithFields(logrus.Fields{
		"shareSessions": u.props.oneSession,
		"sessionId":     u.props.sessionId,
		"authToken":     u.props.authToken,
		"signature":     signature,
	})
	ctxLog.Trace("UserApi.CopyStrategy")

	body := strategies.CopyRequest{SourceStrategySignature: signature}
	err = subAndParse(prepPost(u.builder.Strategies(), u.props).
		MarshalBody(body, marshaler), &res)

	if err != nil {
		ctxLog.WithField("error", err).Debug("UserApi.CopyStrategy request failed")
	}

	return
}

func (u *userApi) MustCopyStrategy(sig string) (res strategies.CopyResponse) {
	logger.WithFields(logrus.Fields{
		"shareSessions": u.props.oneSession,
		"sessionId":     u.props.sessionId,
		"authToken":     u.props.authToken,
		"signature":     sig,
	}).Trace("Api.MustCopyStrategy")

	out, err := u.CopyStrategy(sig)
	if err != nil {
		logger.Panic(err)
	}

	return out
}

func (u *userApi) GetStrategy(id uint) (res strategies.Strategy, err error) {
	ctxLog := logger.WithFields(logrus.Fields{
		"shareSessions": u.props.oneSession,
		"sessionId":     u.props.sessionId,
		"authToken":     u.props.authToken,
		"strategyId":    id,
	})
	ctxLog.Trace("UserApi.GetStrategy")

	err = subAndParse(prepGet(u.builder.Strategy(id), u.props), &res)

	if err != nil {
		ctxLog.WithField("error", err).Debug("UserApi.GetStrategy request failed")
	}

	return
}

func (u *userApi) MustGetStrategy(id uint) (res strategies.Strategy) {
	logger.WithFields(logrus.Fields{
		"shareSessions": u.props.oneSession,
		"sessionId":     u.props.sessionId,
		"authToken":     u.props.authToken,
		"strategyId":    id,
	}).Trace("Api.MustGetStrategy")

	out, err := u.GetStrategy(id)
	if err != nil {
		logger.Panic(err)
	}

	return out
}
