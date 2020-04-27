package strategy

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/step"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
)

type CreateRequest struct {
	Description optional.String `json:"description,omitempty"`
	Name        string          `json:"name"`
	SavedName   optional.String `json:"savedName,omitempty"`
	IsSaved     bool            `json:"isSaved"`
	IsPublic    bool            `json:"isPublic"`
	StepTree    *step.TreeStep  `json:"stepTree"`
}
