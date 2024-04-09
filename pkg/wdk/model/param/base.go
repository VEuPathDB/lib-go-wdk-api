package param

import "lib-go-wdk-api/pkg/wdk/model/common"

type Base struct {
	Name            string   `json:"name"`
	DisplayName     string   `json:"displayName"`
	Help            string   `json:"help"`
	Kind            Kind     `json:"type"`
	IsVisible       bool     `json:"isVisible"`
	Group           string   `json:"group"`
	IsReadOnly      bool     `json:"isReadOnly"`
	AllowEmptyValue bool     `json:"allowEmptyValue"`
	VisibleHelp     string   `json:"visibleHelp"`
	DependentParams []string `json:"dependentParams"`

	InitialDisplayValue *string `json:"initialDisplayValue"`

	PropertyLists common.Properties `json:"properties"`
}
