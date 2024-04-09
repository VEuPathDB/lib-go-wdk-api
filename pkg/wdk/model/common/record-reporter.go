package common

type RecordReporter struct {
	IsInReport  bool    `json:"isInReport"`
	DisplayName string  `json:"displayName"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Scopes      []Scope `json:"scopes"`
	Type        string  `json:"type"`
}
