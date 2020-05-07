package strategy

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/x/xtime"
)

type List []ShortStrategy

type ShortStrategy struct {
	StrategyId                uint      `json:"strategyId"`
	Description               string    `json:"description"`
	Name                      string    `json:"name"`
	Author                    string    `json:"author"`
	RootStepId                uint      `json:"rootStepId"`
	RecordClassName           *string   `json:"recordClassName"`
	Signature                 string    `json:"signature"`
	CreatedTime               xtime.Time `json:"createdTime"`
	LastViewed                xtime.Time `json:"lastViewed"`
	LastModified              xtime.Time `json:"lastModified"`
	IsPublic                  bool      `json:"isPublic"`
	IsSaved                   bool      `json:"isSaved"`
	IsDeleted                 bool      `json:"isDeleted"`
	IsValid                   bool      `json:"isValid"`
	Organization              string    `json:"organization"`
	EstimatedSize             *int      `json:"estimatedSize"`
	NameOfFirstStep           *string   `json:"nameOfFirstStep"`
	LeafAndTransformStepCount *uint     `json:"leafAndTransformStepCount"`
}
