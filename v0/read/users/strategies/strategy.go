package strategies

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/read/common"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/read/users/steps"
)

type Strategy struct {
	common.StrategyNoStepTree

	StepTree Step `json:"stepTree"`
	Steps map[uint]steps.Step
}
