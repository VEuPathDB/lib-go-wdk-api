package step

import (
	"encoding/json"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/answer"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/validation"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/x/xtime"
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
	CreatedTime             xtime.Time         `json:"createdTime"`
	LastRunTime             xtime.Time         `json:"lastRunTime"`
	DisplayPrefs            json.RawMessage   `json:"displayPreferences"`
}
