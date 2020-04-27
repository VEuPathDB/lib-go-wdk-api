package common

import (
	"encoding/json"
	"fmt"
)

const (
	errBadAllMark = `unrecognized 'all' marker "%s"`
)

type AllMarker string

const (
	AllAttributes AllMarker = "__ALL_ATTRIBUTES__"
	AllTables     AllMarker = "__ALL_TABLES__"
)

func (a AllMarker) IsValid() bool {
	return a == AllAttributes || a == AllTables
}

func (a *AllMarker) UnmarshalJSON(bytes []byte) error {
	var tmp string
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	mark := AllMarker(tmp)

	if !mark.IsValid() {
		return fmt.Errorf(errBadAllMark, tmp)
	}

	*a = mark
	return nil
}

func (a AllMarker) MarshalJSON() ([]byte, error) {
	if !a.IsValid() {
		return nil, fmt.Errorf(errBadAllMark, a)
	}
	return json.Marshal(string(a))
}
