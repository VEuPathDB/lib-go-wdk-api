package wdk

import (
	"fmt"
	"strconv"
)

const (
	urlUserStrategies = "/strategies"
	urlUserStrategy   = urlUserStrategies + "/%d"
)

type UserPathBuilder interface {
	Strategies() string
	Strategy(strategyId uint) string
}

func NewUserPathBuilder(url *ApiUrl, userId uint) UserPathBuilder {
	tmp := *url
	tmp.base = tmp.base + fmt.Sprintf(urlUser, strconv.FormatUint(uint64(userId), 10))
	return &userPathBuilder{&tmp}
}

func CurrentUserPathBuilder(url *ApiUrl) UserPathBuilder {
	tmp := *url
	tmp.base = tmp.base + fmt.Sprintf(urlUser, userCurrent)
	return &userPathBuilder{&tmp}
}

type userPathBuilder struct {
	url *ApiUrl
}

func (u *userPathBuilder) Strategies() string {
	return u.url.base + urlUserStrategies + u.url.query
}

func (u *userPathBuilder) Strategy(stratId uint) string {
	return u.url.base + fmt.Sprintf(urlUserStrategy, stratId) + u.url.query
}
