package step

import (
	"encoding/json"

	"lib-go-wdk-api/internal/xtime"
	"lib-go-wdk-api/pkg/wdk/model/answer"
	"lib-go-wdk-api/pkg/wdk/model/validation"
)

type ShortStep struct {
	Id                      uint64            `json:"id"`
	DisplayName             string            `json:"displayName"`
	ShortDisplayName        string            `json:"shortDisplayName"`
	CustomName              *string           `json:"customName"`
	BaseCustomName          *string           `json:"baseCustomName"`
	IsExpanded              bool              `json:"isExpanded"`
	ExpandedName            *string           `json:"expandedName"`
	IsFiltered              bool              `json:"isFiltered"`
	Description             *string           `json:"description"`
	OwnerId                 uint64            `json:"ownerId"`
	StrategyId              *uint64           `json:"strategyId"`
	HasCompleteStepAnalyses bool              `json:"hasCompleteStepAnalyses"`
	RecordClassName         *string           `json:"recordClassName"`
	SearchName              string            `json:"searchName"`
	SearchConfig            answer.Spec       `json:"searchConfig"`
	ValidationBundle        validation.Bundle `json:"validationBundle"`
	CreatedTime             xtime.DateTime    `json:"createdTime"`
	LastRunTime             xtime.DateTime    `json:"lastRunTime"`
	DisplayPrefs            json.RawMessage   `json:"displayPreferences"`
}
