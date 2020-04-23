package common

import "github.com/VEuPathDB/lib-go-wdk-api/v0/optional"

// RecordTable
type RecordTable struct {
	/* Required Fields */

	Name       string            `json:"name"`
	Attributes []RecordAttribute `json:"attributes"`

	/* Optional Fields */

	DisplayName   optional.String    `json:"displayName"`
	Help          optional.String    `json:"help"`
	IsDisplayable optional.Bool      `json:"isDisplayable"`
	IsInReport    optional.Bool      `json:"isInReport"`
	Properties    OptionalProperties `json:"properties"`
	Type          optional.String    `json:"type"`
}
