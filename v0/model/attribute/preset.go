package attribute

import (
	"encoding/json"
	"fmt"
)

const (
	errBadAllMark = `unrecognized 'all' marker "%s"`
)

type Preset string

const (
	PresetAllAttributes     Preset = "__ALL_ATTRIBUTES__"
	PresetDefaultAttributes Preset = "__DEFAULT_ATTRIBUTES__"
)

func (a Preset) Ptr() *Preset {
	return &a
}

func (a Preset) IsValid() bool {
	return a == PresetAllAttributes || a == PresetAllTables
}

func (a *Preset) UnmarshalJSON(bytes []byte) error {
	var tmp string
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	mark := Preset(tmp)

	if !mark.IsValid() {
		return fmt.Errorf(errBadAllMark, tmp)
	}

	*a = mark
	return nil
}

func (a Preset) MarshalJSON() ([]byte, error) {
	if !a.IsValid() {
		return nil, fmt.Errorf(errBadAllMark, a)
	}
	return json.Marshal(string(a))
}
