package param

import "github.com/VEuPathDB/lib-go-wdk-api/v0/service/common"

type Base struct {
	/* Required Fields */

	DependentParams []string `json:"dependentParams"`
	DisplayName     string   `json:"displayName"`
	Group           string   `json:"group"`
	Help            string   `json:"help"`
	IsReadOnly      bool     `json:"isReadOnly"`
	IsVisible       bool     `json:"isVisible"`
	Name            string   `json:"name"`

	/* Optional Fields */

	InitialDisplayValue common.OptionalString `json:"initialDisplayValue"`

	// TODO: Does this belong in the base param?  the schema for this is borked
	Length common.OptionalUint `json:"length"`
}
