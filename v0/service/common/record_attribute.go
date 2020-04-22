package common

// RecordAttribute
type RecordAttribute struct {
	/* Required Properties */

	Name string `json:"name"`

	/* Optional Properties */

	Align         OptionalAlign       `json:"align"`
	DisplayName   OptionalString      `json:"displayName"`
	Formats       RecordReporterArray `json:"formats"`
	Help          OptionalString      `json:"help"`
	IsDisplayable OptionalBool        `json:"isDisplayable"`
	IsInReport    OptionalBool        `json:"isInReport"`
	IsRemovable   OptionalBool        `json:"isRemovable"`
	IsSortable    OptionalBool        `json:"isSortable"`
	Properties    OptionalProperties  `json:"properties"`
	TruncateTo    OptionalUint        `json:"truncateTo"`
	Type          OptionalString      `json:"type"`
}
