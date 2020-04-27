package strategy

import "github.com/VEuPathDB/lib-go-wdk-api/v0/model/step"

type FullStrategy struct {
	ShortStrategy

	StepTree *step.TreeStep           `json:"stepTree"`
	Steps    map[uint64]step.FullStep `json:"steps"`
}
