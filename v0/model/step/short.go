package step

import (
	"encoding/json"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/answer"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/validation"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
	"time"
)

type ShortStep struct {
	Id                      uint64            `json:"id"`
	DisplayName             string            `json:"displayName"`
	ShortDisplayName        string            `json:"shortDisplayName"`
	CustomName              optional.String   `json:"customName"`
	BaseCustomName          optional.String   `json:"baseCustomName"`
	IsExpanded              bool              `json:"isExpanded"`
	ExpandedName            optional.String   `json:"expandedName"`
	IsFiltered              bool              `json:"isFiltered"`
	Description             optional.String   `json:"description"`
	OwnerId                 uint64            `json:"ownerId"`
	StrategyId              optional.Uint64   `json:"strategyId"`
	HasCompleteStepAnalyses bool              `json:"hasCompleteStepAnalyses"`
	RecordClassName         optional.String   `json:"recordClassName"`
	SearchName              string            `json:"searchName"`
	SearchConfig            answer.Spec       `json:"searchConfig"`
	ValidationBundle        validation.Bundle `json:"validationBundle"`
	CreatedTime             time.Time         `json:"createdTime"`
	LastRunTime             time.Time         `json:"lastRunTime"`
	DisplayPrefs            json.RawMessage   `json:"displayPreferences"`
}
