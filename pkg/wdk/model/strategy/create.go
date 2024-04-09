package strategy

import "lib-go-wdk-api/pkg/wdk/model/step"

type CreateRequest struct {
	Description *string        `json:"description,omitempty"`
	Name        string         `json:"name"`
	SavedName   *string        `json:"savedName,omitempty"`
	IsSaved     bool           `json:"isSaved"`
	IsPublic    bool           `json:"isPublic"`
	StepTree    *step.TreeStep `json:"stepTree"`
}
