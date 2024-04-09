package param

type Number struct {
	Base

	Min       float64 `json:"min"`
	Max       float64 `json:"max"`
	Increment float64 `json:"increment"`
}

type NumberRange struct {
	Base

	MinValue  float64 `json:"minValue"`
	MaxValue  float64 `json:"maxValue"`
	Increment float64 `json:"increment"`
}
