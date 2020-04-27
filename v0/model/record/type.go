package record

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/answer"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/attribute"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/common"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/search"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
)

type Type struct {
	FullName                     string               `json:"fullName"`
	DisplayName                  optional.String      `json:"displayName"`
	DisplayNamePlural            optional.String      `json:"displayNamePlural"`
	ShortDisplayName             optional.String      `json:"shortDisplayName"`
	ShortDisplayNamePlural       optional.String      `json:"shortDisplayNamePlural"`
	NativeDisplayName            optional.String      `json:"nativeDisplayName"`
	NativeDisplayNamePlural      optional.String      `json:"nativeDisplayNamePlural"`
	NativeShortDisplayName       optional.String      `json:"nativeShortDisplayName"`
	NativeShortDisplayNamePlural optional.String      `json:"nativeShortDisplayNamePlural"`
	UrlSegment                   string               `json:"urlSegment"`
	IconName                     optional.String      `json:"iconName"`
	UseBasket                    bool                 `json:"useBasket"`
	Description                  optional.String      `json:"description"`
	Properties                   common.Properties    `json:"properties"`
	Formats                      []answer.Format      `json:"formats"`
	HasAllRecordsQuery           bool                 `json:"hasAllRecordsQuery"`
	PrimaryKeyRefs               []string             `json:"primaryKeyColumnRefs"`
	Attributes                   []attribute.Field    `json:"attributes"`
	Tables                       []Table              `json:"tables"`
	Searches                     []search.ShortSearch `json:"searches"`
}
