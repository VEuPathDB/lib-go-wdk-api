package param

import (
	"encoding/json"
	"fmt"
)

// Date
//
// See org.gusdb.wdk.service.formatter.param.DateParamFormatter
type Date struct {
	Base

	MinDate string `json:"minDate"`
	MaxDate string `json:"maxDate"`
}

type aliasDate Date

func (n Date) MarshalJSON() ([]byte, error) {
	if n.Kind != KindDate {
		return nil, fmt.Errorf("invalid type \"%s\", should be \"%s\"", n.Kind, KindDate)
	}

	return json.Marshal(struct {aliasDate}{(aliasDate)(n)})
}

