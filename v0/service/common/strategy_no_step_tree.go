package common

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
	"time"
)

type StrategyNoStepTree struct {
	StrategyId                uint            `json:"strategyId"`
	Description               optional.String `json:"description"`
	Name                      optional.String `json:"name"`
	Author                    optional.String `json:"author"`
	RootStepId                uint            `json:"rootStepId"`
	RecordClassName           optional.String `json:"recordClassName"`
	Signature                 optional.String `json:"signature"`
	CreatedTime               optional.String `json:"createdTime"`
	LastViewed                optional.String `json:"lastViewed"`
	LastModified              optional.String `json:"lastModified"`
	ReleaseVersion            optional.String `json:"releaseVersion"`
	IsPublic                  bool            `json:"isPublic"`
	IsSaved                   bool            `json:"isSaved"`
	IsDeleted                 bool            `json:"isDeleted"`
	IsValid                   bool            `json:"isValid"`
	Organization              optional.String `json:"organization"`
	EstimatedSize             optional.Int            `json:"estimatedSize"`
	LeafAndTransformStepCount optional.Uint   `json:"leafAndTransformStepCount"`
	NameOfFirstStep           optional.String `json:"nameOfFirstStep"`
}

func (l *StrategyNoStepTree) ParseCreatedTime() time.Time {
	if out, err := time.Parse(time.RFC3339Nano, l.CreatedTime.Get()); err != nil {
		panic(err)
	} else {
		return out
	}
}

func (l *StrategyNoStepTree) ParseLastViewed() time.Time {
	if out, err := time.Parse(time.RFC3339Nano, l.LastViewed.Get()); err != nil {
		panic(err)
	} else {
		return out
	}
}

func (l *StrategyNoStepTree) ParseLastModified() time.Time {
	if out, err := time.Parse(time.RFC3339Nano, l.LastModified.Get()); err != nil {
		panic(err)
	} else {
		return out
	}
}
