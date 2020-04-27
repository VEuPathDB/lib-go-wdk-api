package answer

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/common"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
)

// Format
//
// See org.gusdb.wdk.service.formatter.RecordClassFormatter#getAnswerFormatsJson
type Format struct {
	Name        string          `json:"name"`
	Type        string          `json:"type"`
	DisplayName string          `json:"displayName"`
	Description optional.String `json:"description"`
	IsInReport  bool            `json:"isInReport"`
	Scopes      []common.Scope  `json:"scopes"`
}
