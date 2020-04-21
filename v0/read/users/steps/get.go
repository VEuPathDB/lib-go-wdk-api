package steps

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/read/answers"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/read/common"
)

type Step struct {
	Id                      uint                      `json:"id"`
	OwnerId                 uint                      `json:"ownerId"`
	DisplayName             *string                   `json:"displayName"`
	ShortDisplayName        *string                   `json:"shortDisplayName"`
	Description             *string                   `json:"description"`
	CustomName              *string                   `json:"customName"`
	BaseCustomName          *string                   `json:"baseCustomName"`
	SearchName              string                    `json:"searchName"`
	SearchConfig            answers.AnswerSpecRequest `json:"searchConfig"`
	Validation              common.ValidationBundle   `json:"validation"`
	Expanded                bool                      `json:"expanded"`
	ExpandedName            *string                   `json:"expandedName"`
	HasCompleteStepAnalyses bool                      `json:"hasCompleteStepAnalyses"`
	IsFiltered              *bool                     `json:"isFiltered"`
	RecordClassName         *string                   `json:"recordClassName"`
	StrategyId              *uint                     `json:"strategyId"`
	EstimatedSize           *uint                     `json:"estimatedSize"`
	CreatedTime             string                    `json:"createdTime"`
	LastRunTime             string                    `json:"lastRunTime"`
	DisplayPreferences      *common.DisplayPrefs      `json:"displayPreferences"`
}

func (s *Step) HasDisplayName() bool {
	return s.DisplayName != nil
}

func (s *Step) HasShortDisplayName() bool {
	return s.ShortDisplayName != nil
}

func (s *Step) HasDescription() bool {
	return s.Description != nil
}

func (s *Step) HasCustomName() bool {
	return s.CustomName != nil
}

func (s *Step) HasBaseCustomName() bool {
	return s.BaseCustomName != nil
}

func (s *Step) HasExpandedName() bool {
	return s.ExpandedName != nil
}

func (s *Step) HasIsFiltered() bool {
	return s.IsFiltered != nil
}

func (s *Step) HasRecordClassName() bool {
	return s.RecordClassName != nil
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
