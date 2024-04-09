package path

type UserPathBuilder interface {
	Strategies() string
	Strategy(strategyId uint) string
}

func CurrentUserPathBuilder(url ApiUrl) UserPathBuilder

func NewUserBuilder(url ApiUrl, userId uint) UserPathBuilder
