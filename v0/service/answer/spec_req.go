package answer

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/common"
)

type SpecRequest struct {
	Parameters       map[string]string                     `json:"parameters"`
	LegacyFilterName *string                               `json:"legacyFilterName,omitempty"`
	WdkWeight        *int                                  `json:"wdkWeight,omitempty"`
	Filters          []common.FilterSpec                   `json:"filters,omitempty"`
	ColumnFilters    common.ColumnFilters `json:"columnFilters,omitempty"`
}

func (s *SpecRequest) AppendColumnFilter(a, b string, val interface{}) error {
	if s.ColumnFilters == nil {
		s.ColumnFilters = make(common.ColumnFilters)
	}

	return s.ColumnFilters.Append(a, b, val)
}
