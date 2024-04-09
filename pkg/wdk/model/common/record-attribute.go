package common

type RecordAttribute struct {
	Name string `json:"name"`

	Align         *Align              `json:"align"`
	DisplayName   *string             `json:"displayName"`
	Formats       RecordReporterArray `json:"formats"`
	Help          *string             `json:"help"`
	IsDisplayable *bool               `json:"isDisplayable"`
	IsInReport    *bool               `json:"isInReport"`
	IsRemovable   *bool               `json:"isRemovable"`
	IsSortable    *bool               `json:"isSortable"`
	Properties    OptionalProperties  `json:"properties"`
	TruncateTo    *uint               `json:"truncateTo"`
	Type          *string             `json:"type"`
}
