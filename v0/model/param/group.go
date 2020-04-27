package param

type Group struct {
	Name        string   `json:"name"`
	DisplayName string   `json:"displayName"`
	Description string   `json:"description"`
	IsVisible   bool     `json:"isVisible"`
	DisplayType string   `json:"displayType"`
	Parameters  []string `json:"parameters"`
}
