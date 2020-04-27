package path

import (
	"fmt"
	"strconv"
)

const (
	userCurrent = "current"

	urlUserStrategies = "/strategies"
	urlUserStrategy   = urlUserStrategies + "/%d"
)

type UserPathBuilder interface {
	Strategies() string
	Strategy(strategyId uint) string
}

func NewUserBuilder(url ApiUrl, userId uint) UserPathBuilder {
	return &userPathBuilder{
		url: newApiUrl(url.BaseUrl() +
			fmt.Sprintf(urlUser, strconv.FormatUint(uint64(userId), 10))),
	}
}

func CurrentPathBuilder(url ApiUrl) UserPathBuilder {
	return &userPathBuilder{
		url: newApiUrl(url.BaseUrl() + fmt.Sprintf(urlUser, userCurrent)),
	}
}

type userPathBuilder struct {
	url ApiUrl
}

func (u *userPathBuilder) Strategies() string {
	return u.url.wrap(urlUserStrategies)
}

func (u *userPathBuilder) Strategy(stratId uint) string {
	return u.url.wrap(fmt.Sprintf(urlUserStrategy, stratId))
}
