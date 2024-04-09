package answer

import "lib-go-wdk-api/pkg/wdk/model/common"

type PostResponse struct {
	Records []common.RecordInstanceResponse `json:"records"`
	Meta    PostResponseMeta                `json:"meta"`
}

type PostResponseMeta struct {
	RecordClassName        *string  `json:"recordClassName"`
	TotalCount             *uint    `json:"totalCount"`
	ResponseCount          *uint    `json:"responseCount"`
	Tables                 []string `json:"tables"`
	Attributes             []string `json:"attributes"`
	CachePreviouslyExisted bool     `json:"cachePreviouslyExisted"`
}

type Spec struct {
	Parameters       map[string]string `json:"parameters"`
	Filters          []Filter          `json:"filters"`
	ColumnFilters    ColumnFilterMap   `json:"columnFilters"`
	WdkWeight        int               `json:"wdkWeight"`
	LegacyFilterName *string           `json:"legacyFilterName"`
}
