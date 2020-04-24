package strategies

import "github.com/VEuPathDB/lib-go-wdk-api/v0/optional"

type CopyRequest struct {
	SourceStrategySignature string `json:"sourceStrategySignature"`
}

type CreateRequest struct {
	Description optional.String `json:"description,omitempty"`
	Name        string          `json:"name"`
	SavedName   optional.String `json:"savedName,omitempty"`
	IsSaved     bool            `json:"isSaved"`
	IsPublic    bool            `json:"isPublic"`
	StepTree    Step            `json:"stepTree"`
}
