package search

// Filter
//
// See org.gusdb.wdk.service.formatter.QuestionFormatter#getFiltersJson
type Filter struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	IsViewOnly  bool   `json:"isViewOnly"`
}
