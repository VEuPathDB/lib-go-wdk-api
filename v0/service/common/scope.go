package common

import (
	"encoding/json"
	"fmt"
)

const (
	errBadScope = `unrecognized scope value "%s"`
)

// Scope
type Scope string

const (
	ScopeRecord Scope = "record"
	ScopeResult Scope = "results"
)

func (s Scope) IsValid() bool {
	return s == ScopeRecord || s == ScopeResult
}

func (s *Scope) UnmarshalJSON(bytes []byte) error {
	var tmp string

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	scope := Scope(tmp)

	if !scope.IsValid() {
		return fmt.Errorf(errBadScope, tmp)
	}

	*s = scope
	return nil
}

func (s Scope) MarshalJSON() ([]byte, error) {
	if !s.IsValid() {
		return nil, fmt.Errorf(errBadScope, s)
	}
	return json.Marshal(string(s))
}
