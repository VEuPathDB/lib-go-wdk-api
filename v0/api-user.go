package wdk

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/path"
	"github.com/sirupsen/logrus"

	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/strategy"
)

type UserApi interface {
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

type userApi struct {
	builder path.UserBuilder
	userId  string
	props   *apiProps
}

func (u *userApi) ctxLog() *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"shareSessions": u.props.oneSession,
		"sessionId":     u.props.sessionId,
		"authToken":     u.props.authToken,
	})
}

func (u *userApi) GetStrategies() (res strategy.List, err error) {
	u.ctxLog().Trace("UserApi.GetStrategies")

	err = subAndParse(prepGet(u.builder.Strategies(), u.props), &res)

	if err != nil {
		u.ctxLog().WithField("error", err).
			Debug("UserApi.GetStrategies request failed")
	}

	return
}

func (u *userApi) MustGetStrategies() (res strategy.List) {
	u.ctxLog().Trace("UserApi.MustGetStrategies")

	out, err := u.GetStrategies()
	if err != nil {
		u.ctxLog().Panic(err)
	}

	return out
}

func (u *userApi) CopyStrategy(signature string) (res strategy.CopyResponse, err error) {
	ctxLog := u.ctxLog().WithField("signature", signature)
	ctxLog.Trace("UserApi.CopyStrategy")

	body := strategy.CopyRequest{SourceStrategySignature: signature}
	err = subAndParse(prepPost(u.builder.Strategies(), u.props).
		MarshalBody(body, marshaler), &res)

	if err != nil {
		ctxLog.WithField("error", err).
			Debug("UserApi.CopyStrategy request failed")
	}

	return
}

func (u *userApi) MustCopyStrategy(sig string) (res strategy.CopyResponse) {
	u.ctxLog().WithField("signature", sig).Trace("Api.MustCopyStrategy")

	out, err := u.CopyStrategy(sig)
	if err != nil {
		u.ctxLog().Panic(err)
	}

	return out
}

func (u *userApi) GetStrategy(id uint) (res strategy.FullStrategy, err error) {
	ctxLog := u.ctxLog().WithField("strategyId", id)
	ctxLog.Trace("UserApi.GetStrategy")

	err = subAndParse(prepGet(u.builder.Strategy(id), u.props), &res)

	if err != nil {
		ctxLog.WithField("error", err).
			Debug("UserApi.GetStrategy request failed")
	}

	return
}

func (u *userApi) MustGetStrategy(id uint) (res strategy.FullStrategy) {
	u.ctxLog().WithField("strategyId", id).Trace("Api.MustGetStrategy")

	out, err := u.GetStrategy(id)
	if err != nil {
		u.ctxLog().Panic(err)
	}

	return out
}
