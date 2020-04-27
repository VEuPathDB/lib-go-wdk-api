package attribute

// ToolSet
//
// See org.gusdb.wdk.service.formatter.AttributeFieldFormatter#formatToolsJson
type ToolSet struct {
	Reports []string `json:"reports"`
	Filters []string `json:"filters"`
}
