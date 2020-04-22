package common

// RecordTable
type RecordTable struct {
	/* Required Fields */

	Name       string            `json:"name"`
	Attributes []RecordAttribute `json:"attributes"`

	/* Optional Fields */

	DisplayName   OptionalString     `json:"displayName"`
	Help          OptionalString     `json:"help"`
	IsDisplayable OptionalBool       `json:"isDisplayable"`
	IsInReport    OptionalBool       `json:"isInReport"`
	Properties    OptionalProperties `json:"properties"`
	Type          OptionalString     `json:"type"`
}
