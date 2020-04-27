package common

import (
	"encoding/json"
	"errors"
	"strings"
)

type AllOrList struct {
	isAll bool
	mark  AllMarker
	attrs []string
}

func (a *AllOrList) UnmarshalJSON(bytes []byte) error {
	var all AllMarker
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

func (a *AllOrList) SetAll(mark AllMarker) {
	a.isAll = true
	a.mark = mark
	a.attrs = nil
}

func (a *AllOrList) SetList(list []string) {
	a.isAll = false
	a.mark = ""
	a.attrs = list
}

func (a *AllOrList) IsAll() bool {
	return a.isAll
}

func (a *AllOrList) IsList() bool {
	return !a.isAll
}

func (a *AllOrList) AllMarker() AllMarker {
	return a.mark
}

func (a *AllOrList) List() []string {
	return a.attrs
}
