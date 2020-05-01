package answer

type Spec struct {
	Parameters       map[string]string `json:"parameters"`
	Filters          []Filter          `json:"filters"`
	ColumnFilters    ColumnFilterMap   `json:"columnFilters"`
	WdkWeight        int               `json:"wdkWeight"`
	LegacyFilterName *string           `json:"legacyFilterName"`
}
