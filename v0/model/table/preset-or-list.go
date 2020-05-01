package table

import (
	"encoding/json"
	"errors"
)

const allMarker = "__ALL_TABLES__"

type PresetOrList struct {
	isAll bool
	attrs []string
}

func (a *PresetOrList) UnmarshalJSON(bytes []byte) error {
	var all string
	if err := json.Unmarshal(bytes, &all); err == nil {
		if all != allMarker {
			return errors.New("invalid marker value " + all)
		}
		a.isAll = true
		return nil
	}

	var attrs []string
	if err := json.Unmarshal(bytes, &attrs); err == nil {
		a.attrs = attrs
		return nil
	}

	return errors.New("unrecognized json, should be string or string[]")
}

func (a *PresetOrList) MarshalJSON() ([]byte, error) {
	if a.isAll {
		return json.Marshal(allMarker)
	}
	return json.Marshal(a.attrs)
}

func (a *PresetOrList) SetAll() {
	a.isAll = true
	a.attrs = nil
}

func (a *PresetOrList) SetList(list []string) {
	a.isAll = false
	a.attrs = list
}

func (a *PresetOrList) IsAll() bool {
	return a.isAll
}

func (a *PresetOrList) IsList() bool {
	return !a.isAll
}

func (a *PresetOrList) List() []string {
	return a.attrs
}
