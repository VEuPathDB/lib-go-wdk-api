package strategy

import "lib-go-wdk-api/pkg/wdk/model/step"

type FullStrategy struct {
	ShortStrategy

	StepTree *step.TreeStep           `json:"stepTree"`
	Steps    map[uint64]step.FullStep `json:"steps"`
}
