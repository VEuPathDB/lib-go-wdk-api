package answer

import "github.com/VEuPathDB/lib-go-wdk-api/v0/optional"

type Spec struct {
	Parameters       map[string]string `json:"parameters"`
	Filters          []Filter          `json:"filters"`
	ColumnFilters    ColumnFilterMap   `json:"columnFilters"`
	WdkWeight        int               `json:"wdkWeight"`
	LegacyFilterName optional.String   `json:"legacyFilterName"`
}
