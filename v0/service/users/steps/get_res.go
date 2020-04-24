package steps

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/answer"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/common"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/validation"
)

type Step struct {
	Id                      uint                 `json:"id"`
	OwnerId                 uint                 `json:"ownerId"`
	DisplayName             optional.String      `json:"displayName"`
	ShortDisplayName        optional.String      `json:"shortDisplayName"`
	Description             optional.String      `json:"description"`
	CustomName              optional.String      `json:"customName"`
	BaseCustomName          optional.String      `json:"baseCustomName"`
	SearchName              string               `json:"searchName"`
	SearchConfig            answer.SpecRequest   `json:"searchConfig"`
	Validation              validation.Bundle    `json:"validation"`
	Expanded                bool                 `json:"expanded"`
	ExpandedName            optional.String      `json:"expandedName"`
	HasCompleteStepAnalyses bool                 `json:"hasCompleteStepAnalyses"`
	IsFiltered              optional.Bool        `json:"isFiltered"`
	RecordClassName         optional.String      `json:"recordClassName"`
	StrategyId              optional.Uint        `json:"strategyId"`
	EstimatedSize           optional.Int         `json:"estimatedSize"`
	CreatedTime             string               `json:"createdTime"`
	LastRunTime             string               `json:"lastRunTime"`
	DisplayPreferences      *common.DisplayPrefs `json:"displayPreferences"`
}
