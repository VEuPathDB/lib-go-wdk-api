package strategy

import (
	"lib-go-wdk-api/internal/xtime"
	"lib-go-wdk-api/pkg/wdk/model/validation"
)

type ShortStrategy struct {
	StrategyId                uint               `json:"strategyId"`
	Description               string             `json:"description"`
	Name                      string             `json:"name"`
	Author                    string             `json:"author"`
	RootStepId                uint               `json:"rootStepId"`
	RecordClassName           *string            `json:"recordClassName"`
	Signature                 string             `json:"signature"`
	CreatedTime               xtime.DateTime     `json:"createdTime"`
	LastViewed                xtime.DateTime     `json:"lastViewed"`
	LastModified              xtime.DateTime     `json:"lastModified"`
	ReleaseVersion            string             `json:"releaseVersion"`
	IsPublic                  bool               `json:"isPublic"`
	IsSaved                   bool               `json:"isSaved"`
	IsDeleted                 bool               `json:"isDeleted"`
	IsValid                   bool               `json:"isValid"`
	Organization              string             `json:"organization"`
	EstimatedSize             *int               `json:"estimatedSize"`
	NameOfFirstStep           *string            `json:"nameOfFirstStep"`
	LeafAndTransformStepCount *uint              `json:"leafAndTransformStepCount"`
	Validation                *validation.Bundle `json:"validation"`
}
