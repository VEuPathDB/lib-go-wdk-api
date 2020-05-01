package attribute

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/answer"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/common"
)

// Field
//
// See org.gusdb.wdk.service.formatter.AttributeFieldFormatter
type Field struct {
	Align         *common.Align     `json:"align"`
	DisplayName   string            `json:"displayName"`
	Formats       []answer.Format   `json:"formats"`
	Help          *string           `json:"help"`
	IsDisplayable bool              `json:"isDisplayable"`
	IsInReport    bool              `json:"isInReport"`
	IsRemovable   bool              `json:"isRemovable"`
	IsSortable    bool              `json:"isSortable"`
	Name          string            `json:"name"`
	Properties    common.Properties `json:"properties"`
	TruncateTo    uint              `json:"truncateTo"`
	Type          string            `json:"type"`
	Tools         ToolSet           `json:"tools"`
	DataType      FieldDataType     `json:"columnDataType"`
}
