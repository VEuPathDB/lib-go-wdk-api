package param

// Base
//
// See org.gusdb.wdk.service.formatter.param.ParamFormatter
type Base struct {

	/* WDK Required Fields */

	Name            string   `json:"name"`
	DisplayName     string   `json:"displayName"`
	Help            string   `json:"help"`
	Kind            Kind     `json:"type"`
	IsVisible       bool     `json:"isVisible"`
	Group           string   `json:"group"`
	IsReadOnly      bool     `json:"isReadOnly"`
	DependentParams []string `json:"dependentParams"`

	/* WDK Optional Fields */

	InitialDisplayValue *string `json:"initialDisplayValue"`
}
