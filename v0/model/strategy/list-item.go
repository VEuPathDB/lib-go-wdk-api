package strategy

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
	"time"
)

type List []ShortStrategy

type ShortStrategy struct {
	StrategyId                uint            `json:"strategyId"`
	Description               string          `json:"description"`
	Name                      string          `json:"name"`
	Author                    string          `json:"author"`
	RootStepId                uint            `json:"rootStepId"`
	RecordClassName           optional.String `json:"recordClassName"`
	Signature                 string          `json:"signature"`
	CreatedTime               time.Time       `json:"createdTime"`
	LastViewed                time.Time       `json:"lastViewed"`
	LastModified              time.Time       `json:"lastModified"`
	IsPublic                  bool            `json:"isPublic"`
	IsSaved                   bool            `json:"isSaved"`
	IsDeleted                 bool            `json:"isDeleted"`
	IsValid                   bool            `json:"isValid"`
	Organization              string          `json:"organization"`
	EstimatedSize             optional.Int    `json:"estimatedSize"`
	NameOfFirstStep           optional.String `json:"nameOfFirstStep"`
	LeafAndTransformStepCount optional.Uint   `json:"leafAndTransformStepCount"`
}
