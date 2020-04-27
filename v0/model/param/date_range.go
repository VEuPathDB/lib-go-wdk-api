package param

import (
	"encoding/json"
	"fmt"
)

type DateRange struct {
	Base

	MinDate string `json:"minDate"`
	MaxDate string `json:"maxDate"`
}

type aliasDateRange DateRange

func (n DateRange) MarshalJSON() ([]byte, error) {
	if n.Kind != KindDateRange {
		return nil, fmt.Errorf("invalid type \"%s\", should be \"%s\"", n.Kind, KindDateRange)
	}

	return json.Marshal(struct{ aliasDateRange }{(aliasDateRange)(n)})
}
