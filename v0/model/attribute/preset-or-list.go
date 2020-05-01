package attribute

import (
	"encoding/json"
	"errors"
	"strings"
)

type PresetOrList struct {
	isAll bool
	mark  Preset
	attrs []string
}

func (a *PresetOrList) UnmarshalJSON(bytes []byte) error {
	var all Preset
	if err := all.UnmarshalJSON(bytes); err != nil {
		if strings.HasPrefix(err.Error(), "unrecognized 'all'") {
			return err
		}
	} else {
		a.isAll = true
		a.mark = all
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
		return json.Marshal(a.mark)
	}
	return json.Marshal(a.attrs)
}

func (a *PresetOrList) SetAll(mark Preset) {
	a.isAll = true
	a.mark = mark
	a.attrs = nil
}

func (a *PresetOrList) SetList(list []string) {
	a.isAll = false
	a.mark = ""
	a.attrs = list
}

func (a *PresetOrList) IsAll() bool {
	return a.isAll
}

func (a *PresetOrList) IsList() bool {
	return !a.isAll
}

func (a *PresetOrList) AllMarker() Preset {
	return a.mark
}

func (a *PresetOrList) List() []string {
	return a.attrs
}
