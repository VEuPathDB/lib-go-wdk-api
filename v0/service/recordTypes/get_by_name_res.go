package recordTypes

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/common"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/recordTypes/searches"
)

type RecordTypeResponse struct {
	/* Required Fields */

	FullName               string                     `json:"fullName"`
	DisplayName            string                     `json:"displayName"`
	DisplayNamePlural      string                     `json:"displayNamePlural"`
	ShortDisplayName       string                     `json:"shortDisplayName"`
	ShortDisplayNamePlural string                     `json:"shortDisplayNamePlural"`
	UrlSegment             string                     `json:"urlSegment"`
	UseBasket              bool                       `json:"useBasket"`
	Formats                common.RecordReporterArray `json:"formats"`
	HasAllRecordsQuery     bool                       `json:"hasAllRecordsQuery"`
	PrimaryKeyColumnRefs   []string                   `json:"primaryKeyColumnRefs"`
	Description            string                     `json:"description"`
	Attributes             []common.RecordAttribute   `json:"attributes"`
	Tables                 []common.RecordTable       `json:"tables"`
	Searches               []searches.SearchResponse  `json:"searches"`

	/* Optional Fields */

	NativeDisplayName            common.OptionalString     `json:"nativeDisplayName"`
	NativeDisplayNamePlural      common.OptionalString     `json:"nativeDisplayNamePlural"`
	NativeShortDisplayName       common.OptionalString     `json:"nativeShortDisplayName"`
	NativeShortDisplayNamePlural common.OptionalString     `json:"nativeShortDisplayNamePlural"`
	IconName                     common.OptionalString     `json:"iconName"`
	Properties                   common.OptionalProperties `json:"properties"`
	RecordIdAttributeName        common.OptionalString     `json:"recordIdAttributeName"`
}
