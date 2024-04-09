package record

import (
	"lib-go-wdk-api/pkg/wdk/model/answer"
	"lib-go-wdk-api/pkg/wdk/model/attribute"
	"lib-go-wdk-api/pkg/wdk/model/common"
	"lib-go-wdk-api/pkg/wdk/model/search"
)

type Type struct {
	FullName                     string               `json:"fullName"`
	DisplayName                  *string              `json:"displayName"`
	DisplayNamePlural            *string              `json:"displayNamePlural"`
	ShortDisplayName             *string              `json:"shortDisplayName"`
	ShortDisplayNamePlural       *string              `json:"shortDisplayNamePlural"`
	NativeDisplayName            *string              `json:"nativeDisplayName"`
	NativeDisplayNamePlural      *string              `json:"nativeDisplayNamePlural"`
	NativeShortDisplayName       *string              `json:"nativeShortDisplayName"`
	NativeShortDisplayNamePlural *string              `json:"nativeShortDisplayNamePlural"`
	UrlSegment                   string               `json:"urlSegment"`
	IconName                     *string              `json:"iconName"`
	UseBasket                    bool                 `json:"useBasket"`
	Description                  *string              `json:"description"`
	Properties                   common.Properties    `json:"properties"`
	Formats                      []answer.Format      `json:"formats"`
	HasAllRecordsQuery           bool                 `json:"hasAllRecordsQuery"`
	PrimaryKeyRefs               []string             `json:"primaryKeyColumnRefs"`
	Attributes                   []attribute.Field    `json:"attributes"`
	Tables                       []Table              `json:"tables"`
	Searches                     []search.ShortSearch `json:"searches"`
}
