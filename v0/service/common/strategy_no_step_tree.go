package common

import "time"

type StrategyNoStepTree struct {
	StrategyId                uint    `json:"strategyId"`
	Description               *string `json:"description"`
	Name                      *string `json:"name"`
	Author                    *string `json:"author"`
	RootStepId                uint    `json:"rootStepId"`
	RecordClassName           *string `json:"recordClassName"`
	Signature                 *string `json:"signature"`
	CreatedTime               *string `json:"createdTime"`
	LastViewed                *string `json:"lastViewed"`
	LastModified              *string `json:"lastModified"`
	ReleaseVersion            *string `json:"releaseVersion"`
	IsPublic                  bool    `json:"isPublic"`
	IsSaved                   bool    `json:"isSaved"`
	IsDeleted                 bool    `json:"isDeleted"`
	IsValid                   bool    `json:"isValid"`
	Organization              *string `json:"organization"`
	EstimatedSize             *int   `json:"estimatedSize"`
	LeafAndTransformStepCount *uint   `json:"leafAndTransformStepCount"`
	NameOfFirstStep           *string `json:"nameOfFirstStep"`
}

func (l *StrategyNoStepTree) HasDescription() bool {
	return l.Description != nil
}

func (l *StrategyNoStepTree) HasName() bool {
	return l.Name != nil
}

func (l *StrategyNoStepTree) HasAuthor() bool {
	return l.Author != nil
}

func (l *StrategyNoStepTree) HasRecordClassName() bool {
	return l.RecordClassName != nil
}

func (l *StrategyNoStepTree) HasSignature() bool {
	return l.Signature != nil
}

func (l *StrategyNoStepTree) HasCreatedTime() bool {
	return l.CreatedTime != nil
}

func (l *StrategyNoStepTree) ParseCreatedTime() time.Time {
	if out, err := time.Parse(time.RFC3339Nano, *l.CreatedTime); err != nil {
		panic(err)
	} else {
		return out
	}
}

func (l *StrategyNoStepTree) HasLastViewed() bool {
	return l.LastViewed != nil
}

func (l *StrategyNoStepTree) ParseLastViewed() time.Time {
	if out, err := time.Parse(time.RFC3339Nano, *l.LastViewed); err != nil {
		panic(err)
	} else {
		return out
	}
}

func (l *StrategyNoStepTree) HasLastModified() bool {
	return l.LastModified != nil
}

func (l *StrategyNoStepTree) ParseLastModified() time.Time {
	if out, err := time.Parse(time.RFC3339Nano, *l.LastModified); err != nil {
		panic(err)
	} else {
		return out
	}
}

func (l *StrategyNoStepTree) HasReleaseVersion() bool {
	return l.ReleaseVersion != nil
}

func (l *StrategyNoStepTree) HasOrganization() bool {
	return l.Organization != nil
}

func (l *StrategyNoStepTree) HasEstimatedSize() bool {
	return l.EstimatedSize != nil
}

func (l *StrategyNoStepTree) HasLeafAndTransformStepCount() bool {
	return l.LeafAndTransformStepCount != nil
}

func (l *StrategyNoStepTree) HasNameOfFirstStep() bool {
	return l.NameOfFirstStep != nil
}
