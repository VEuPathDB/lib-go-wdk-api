package common

import "github.com/VEuPathDB/lib-go-wdk-api/v0/optional"

// RecordAttribute
type RecordAttribute struct {
	/* Required Properties */

	Name string `json:"name"`

	/* Optional Properties */

	Align         OptionalAlign       `json:"align"`
	DisplayName   optional.String     `json:"displayName"`
	Formats       RecordReporterArray `json:"formats"`
	Help          optional.String     `json:"help"`
	IsDisplayable optional.Bool       `json:"isDisplayable"`
	IsInReport    optional.Bool       `json:"isInReport"`
	IsRemovable   optional.Bool       `json:"isRemovable"`
	IsSortable    optional.Bool       `json:"isSortable"`
	Properties    OptionalProperties  `json:"properties"`
	TruncateTo    optional.Uint       `json:"truncateTo"`
	Type          optional.String     `json:"type"`
}
