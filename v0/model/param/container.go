package param

type BasicContainer struct {
	Groups     []Group  `json:"groups"`
	ParamNames []string `json:"paramNames"`
}

type Container struct {
	BasicContainer

	Parameters []MultiParam `json:"parameters"`
}
