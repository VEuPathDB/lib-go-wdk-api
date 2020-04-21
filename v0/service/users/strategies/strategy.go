package strategies

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/common"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/users/steps"
)

type Strategy struct {
	common.StrategyNoStepTree

	StepTree Step `json:"stepTree"`
	Steps    map[uint]steps.Step
}
