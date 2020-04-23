package param

import (
	"encoding/json"
	"fmt"
)

type NumberRange struct {
	Base

	MinValue  float64 `json:"minValue"`
	MaxValue  float64 `json:"maxValue"`
	Increment float64 `json:"increment"`
}

type aliasNumberRange NumberRange

func (n NumberRange) MarshalJSON() ([]byte, error) {
	if n.Kind != KindNumberRange {
		return nil, fmt.Errorf("invalid type \"%s\", should be \"%s\"", n.Kind, KindNumberRange)
	}

	return json.Marshal(struct {aliasNumberRange}{(aliasNumberRange)(n)})
}
