package record

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/attribute"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/common"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
)

type Table struct {
	Attributes     attribute.Field   `json:"attributes"`
	ClientSortSpec []attribute.Sort  `json:"clientSortSpec"`
	Description    optional.String   `json:"description"`
	DisplayName    optional.String   `json:"displayName"`
	Help           optional.String   `json:"help"`
	IsDisplayable  bool              `json:"isDisplayable"`
	IsInReport     bool              `json:"isInReport"`
	Name           optional.String   `json:"name"`
	Properties     common.Properties `json:"properties"`
	Type           string            `json:"type"`
}
