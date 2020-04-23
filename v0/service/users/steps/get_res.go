package steps

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/answer"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/common"
)

type Step struct {
	Id                      uint                    `json:"id"`
	OwnerId                 uint                    `json:"ownerId"`
	DisplayName             optional.String         `json:"displayName"`
	ShortDisplayName        optional.String         `json:"shortDisplayName"`
	Description             optional.String         `json:"description"`
	CustomName              optional.String         `json:"customName"`
	BaseCustomName          optional.String         `json:"baseCustomName"`
	SearchName              string                  `json:"searchName"`
	SearchConfig            answer.SpecRequest      `json:"searchConfig"`
	Validation              common.ValidationBundle `json:"validation"`
	Expanded                bool                    `json:"expanded"`
	ExpandedName            optional.String         `json:"expandedName"`
	HasCompleteStepAnalyses bool                    `json:"hasCompleteStepAnalyses"`
	IsFiltered              *bool                   `json:"isFiltered"`
	RecordClassName         optional.String         `json:"recordClassName"`
	StrategyId              *uint                   `json:"strategyId"`
	EstimatedSize           *int                    `json:"estimatedSize"`
	CreatedTime             string                  `json:"createdTime"`
	LastRunTime             string                  `json:"lastRunTime"`
	DisplayPreferences      *common.DisplayPrefs    `json:"displayPreferences"`
}

func (s *Step) HasIsFiltered() bool {
	return s.IsFiltered != nil
}

func (s *Step) HasStrategyId() bool {
	return s.StrategyId != nil
}

func (s *Step) HasEstimatedSize() bool {
	return s.EstimatedSize != nil
}

func (s *Step) HasDisplayPreferences() bool {
	return s.DisplayPreferences != nil
}
