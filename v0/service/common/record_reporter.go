package common

import (
	"encoding/json"
	"fmt"
)

// Scope
type Scope string

const (
	ScopeRecord Scope = "record"
	ScopeResult Scope = "result"
)

func (s *Scope) UnmarshalJSON(bytes []byte) error {
	var tmp string

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	scope := Scope(tmp)

	switch scope {
	case ScopeRecord:
		fallthrough
	case ScopeResult:
		*s = scope
		return nil
	}

	return fmt.Errorf("unrecognized scope name \"%s\"", tmp)
}

// RecordReporter
type RecordReporter struct {
	IsInReport  bool    `json:"isInReport"`
	DisplayName string  `json:"displayName"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Scopes      []Scope `json:"scopes"`
	Type        string  `json:"type"`
}
