package strategy

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/step"
)

type CreateRequest struct {
	Description *string        `json:"description,omitempty"`
	Name        string         `json:"name"`
	SavedName   *string        `json:"savedName,omitempty"`
	IsSaved     bool           `json:"isSaved"`
	IsPublic    bool           `json:"isPublic"`
	StepTree    *step.TreeStep `json:"stepTree"`
}
