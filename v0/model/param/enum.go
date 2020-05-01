package param

type Enum struct {
	Base

	DisplayType      DisplayType `json:"displayType"`
	MaxSelectedCount uint        `json:"maxSelectedCount"`
	MinSelectedCount uint        `json:"minSelectedCount"`
	Vocabulary       Vocabulary  `json:"vocabulary"`

	/* Tree Box Param Fields */
	CountOnlyLeaves *bool `json:"countOnlyLeaves"`
	DepthExpanded   *uint `json:"depthExpanded"`
}
