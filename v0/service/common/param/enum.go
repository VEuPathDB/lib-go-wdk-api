package param

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
)

type Enum struct {
	Base

	DisplayType      DisplayType `json:"displayType"`
	MaxSelectedCount uint        `json:"maxSelectedCount"`
	MinSelectedCount uint        `json:"minSelectedCount"`
	Vocabulary       Vocabulary  `json:"vocabulary"`

	/* Tree Box Param Fields */
	CountOnlyLeaves optional.Bool `json:"countOnlyLeaves"`
	DepthExpanded   optional.Uint `json:"depthExpanded"`
}
