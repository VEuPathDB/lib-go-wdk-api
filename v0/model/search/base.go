package search

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/attribute"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/common"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/summary"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
)

// Base contains fields common to the different search
// result types.
type Base struct {
	UrlSegment                            string            `json:"urlSegment"`
	FullName                              string            `json:"fullName"`
	DisplayName                           string            `json:"displayName"`
	ShortDisplayName                      string            `json:"shortDisplayName"`
	Description                           optional.String   `json:"description"`
	IconName                              optional.String   `json:"iconName"`
	Summary                               optional.String   `json:"summary"`
	Help                                  optional.String   `json:"help"`
	NewBuild                              optional.String   `json:"newBuild"`
	ReviseBuild                           optional.String   `json:"reviseBuild"`
	OutputRecordClassName                 string            `json:"outputRecordClassName"`
	Filters                               []string          `json:"filters"`
	DefaultAttributes                     []string          `json:"defaultAttributes"`
	DefaultSorting                        []attribute.Sort  `json:"defaultSorting"`
	DynamicAttributes                     []attribute.Field `json:"dynamicAttributes"`
	DefaultSummaryView                    string            `json:"defaultSummaryView"`
	SummaryViewPlugins                    []summary.View    `json:"summaryViewPlugins"`
	IsAnalyzable                          bool              `json:"isAnalyzable"`
	AllowedPrimaryInputRecordClassNames   []string          `json:"allowedPrimaryInputRecordClassNames"`
	AllowedSecondaryInputRecordClassNames []string          `json:"allowedSecondaryInputRecordClassNames"`
	QueryName                             optional.String   `json:"queryName"`
	NoSummaryOnSingleRecord               bool              `json:"noSummaryOnSingleRecord"`
	Properties                            common.Properties `json:"properties"`
}
