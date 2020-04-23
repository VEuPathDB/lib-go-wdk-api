package common

import "time"

type StrategyListingItem struct {
	Description     string  `json:"description"`
	Name            string  `json:"name"`
	Author          string  `json:"author"`
	RecordClassName *string `json:"recordClassName"`
	Signature       string  `json:"signature"`
	Organization    string  `json:"organization"`
	CreatedTime     string  `json:"createdTime"`
	LastViewed      string  `json:"lastViewed"`
	LastModified    string  `json:"lastModified"`
	IsPublic        bool    `json:"isPublic"`
	IsSaved         bool    `json:"isSaved"`
	IsDeleted       bool    `json:"isDeleted"`
	IsValid         bool    `json:"isValid"`
	StrategyId      uint    `json:"strategyId"`
	RootStepId      uint    `json:"rootStepId"`
	EstimatedSize   int     `json:"estimatedSize"`
}

func (l *StrategyListingItem) ParseCreatedTime() time.Time {
	if out, err := time.Parse(time.RFC3339Nano, l.CreatedTime); err != nil {
		panic(err)
	} else {
		return out
	}
}

func (l *StrategyListingItem) ParseLastViewed() time.Time {
	if out, err := time.Parse(time.RFC3339Nano, l.LastViewed); err != nil {
		panic(err)
	} else {
		return out
	}
}

func (l *StrategyListingItem) ParseLastModified() time.Time {
	if out, err := time.Parse(time.RFC3339Nano, l.LastModified); err != nil {
		panic(err)
	} else {
		return out
	}
}
