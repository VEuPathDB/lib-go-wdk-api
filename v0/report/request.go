package report

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/answer"
)

type Request struct {
	SearchConfig answer.Spec      `json:"searchConfig"`
	ReportConfig Config           `json:"reportConfig"`
	ViewFilters  []ViewFilterPair `json:"viewFilters,omitempty"`
}

type ViewFilterPair struct {
	Key      string      `json:"name"`
	Value    interface{} `json:"value"`
	Disabled *bool       `json:"disabled"`
}
