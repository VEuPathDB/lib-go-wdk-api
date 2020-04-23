package param

import (
	"encoding/json"
	"fmt"
)

// Number
//
// See org.gusdb.wdk.service.formatter.param.NumberParamFormatter
type Number struct {
	Base

	Min       float64 `json:"min"`
	Max       float64 `json:"max"`
	Increment float64 `json:"increment"`
}

type aliasNumber Number

func (n Number) MarshalJSON() ([]byte, error) {
	if n.Kind != KindNumber {
		return nil, fmt.Errorf("invalid type \"%s\", should be \"%s\"", n.Kind, KindNumber)
	}

	return json.Marshal(struct {aliasNumber}{(aliasNumber)(n)})
}

