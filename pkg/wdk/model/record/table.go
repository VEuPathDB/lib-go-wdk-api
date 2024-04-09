package record

import (
	"lib-go-wdk-api/pkg/wdk/model/attribute"
	"lib-go-wdk-api/pkg/wdk/model/common"
)

type Table struct {
	Attributes     []attribute.Field `json:"attributes"`
	ClientSortSpec []attribute.Sort  `json:"clientSortSpec"`
	Description    *string           `json:"description"`
	DisplayName    *string           `json:"displayName"`
	Help           *string           `json:"help"`
	IsDisplayable  bool              `json:"isDisplayable"`
	IsInReport     bool              `json:"isInReport"`
	Name           *string           `json:"name"`
	Properties     common.Properties `json:"properties"`
	Type           string            `json:"type"`
}
