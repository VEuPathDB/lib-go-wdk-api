package answers

import (
	"encoding/json"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/read/common"
)

type AnswerSpecRequest struct {
	Parameters       map[string]string   `json:"parameters"`
	LegacyFilterName string              `json:"legacyFilterName"`
	WdkWeight        int                 `json:"wdkWeight"`
	Filters          []common.FilterSpec `json:"filters"`

	// TODO: this can be more specific...
	ColumnFilters map[string]map[string]json.RawMessage `json:"columnFilters"`
}
