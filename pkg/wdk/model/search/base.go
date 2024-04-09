package search

import (
	"lib-go-wdk-api/pkg/wdk/model/attribute"
	"lib-go-wdk-api/pkg/wdk/model/common"
	"lib-go-wdk-api/pkg/wdk/model/summary"
)

// Base contains fields common to the different search result types.
type Base struct {
	UrlSegment                            string            `json:"urlSegment"`
	FullName                              string            `json:"fullName"`
	DisplayName                           string            `json:"displayName"`
	ShortDisplayName                      string            `json:"shortDisplayName"`
	Description                           *string           `json:"description"`
	IconName                              *string           `json:"iconName"`
	Summary                               *string           `json:"summary"`
	Help                                  *string           `json:"help"`
	NewBuild                              *string           `json:"newBuild"`
	ReviseBuild                           *string           `json:"reviseBuild"`
	OutputRecordClassName                 string            `json:"outputRecordClassName"`
	Filters                               []Filter          `json:"filters"`
	DefaultAttributes                     []string          `json:"defaultAttributes"`
	DefaultSorting                        []attribute.Sort  `json:"defaultSorting"`
	DynamicAttributes                     []attribute.Field `json:"dynamicAttributes"`
	DefaultSummaryView                    string            `json:"defaultSummaryView"`
	SummaryViewPlugins                    []summary.View    `json:"summaryViewPlugins"`
	IsAnalyzable                          bool              `json:"isAnalyzable"`
	AllowedPrimaryInputRecordClassNames   []string          `json:"allowedPrimaryInputRecordClassNames"`
	AllowedSecondaryInputRecordClassNames []string          `json:"allowedSecondaryInputRecordClassNames"`
	QueryName                             *string           `json:"queryName"`
	NoSummaryOnSingleRecord               bool              `json:"noSummaryOnSingleRecord"`
	Properties                            common.Properties `json:"properties"`
}
