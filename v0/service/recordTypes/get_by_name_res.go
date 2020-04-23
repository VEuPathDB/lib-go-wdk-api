package recordTypes

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
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

	NativeDisplayName            optional.String           `json:"nativeDisplayName"`
	NativeDisplayNamePlural      optional.String           `json:"nativeDisplayNamePlural"`
	NativeShortDisplayName       optional.String           `json:"nativeShortDisplayName"`
	NativeShortDisplayNamePlural optional.String           `json:"nativeShortDisplayNamePlural"`
	IconName                     optional.String           `json:"iconName"`
	Properties                   common.OptionalProperties `json:"properties"`
	RecordIdAttributeName        optional.String           `json:"recordIdAttributeName"`
}
