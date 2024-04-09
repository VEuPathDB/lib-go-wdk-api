package param

type BasicContainer struct {
	Groups     []Group  `json:"groups"`
	ParamNames []string `json:"paramNames"`
}

type Container struct {
	BasicContainer

	Parameters []MultiParam `json:"parameters"`
}

type Group struct {
	Name        string   `json:"name"`
	DisplayName string   `json:"displayName"`
	Description string   `json:"description"`
	IsVisible   bool     `json:"isVisible"`
	DisplayType string   `json:"displayType"`
	Parameters  []string `json:"parameters"`
}
