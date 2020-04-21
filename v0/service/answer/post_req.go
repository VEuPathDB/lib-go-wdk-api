package answer

import "github.com/VEuPathDB/lib-go-wdk-api/v0/service/common"

type PostRequest struct {
	SearchConfig SpecRequest               `json:"searchConfig"`
	ReportConfig common.AnswerFormatConfig `json:"reportConfig"`
	ViewFilters  []common.FilterSpec       `json:"viewFilters,omitempty"`
}
