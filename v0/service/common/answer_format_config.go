package common

import (
	"encoding/json"
	"errors"
	"fmt"
)

type AnswerFormatConfig struct {
	Pagination struct {
		Offset     uint `json:"offset"`
		NumRecords uint `json:"numRecords"`
	} `json:"pagination"`

	Attributes AllOrList     `json:"attributes"`
	Tables     AllOrList     `json:"tables"`
	Sorting    AttributeSort `json:"sorting"`
}

type AllOrList struct {
	isAll bool
	attrs []string
}

func (a *AllOrList) UnmarshalJSON(bytes []byte) error {
	var all string
	if err := json.Unmarshal(bytes, &all); err == nil {
		switch all {
		case "__ALL_ATTRIBUTES__":
			fallthrough
		case "__ALL_TABLES__":
			a.isAll = true
			return nil
		default:
			return fmt.Errorf("unrecognized value \"%s\"", all)
		}
	}

	var attrs []string
	if err := json.Unmarshal(bytes, &attrs); err == nil {
		a.attrs = attrs
		return nil
	}

	return errors.New("unrecognized json, should be string or string[]")
}

func (a *AllOrList) IsAll() bool {
	return a.isAll
}

func (a *AllOrList) IsList() bool {
	return !a.isAll
}

func (a *AllOrList) List() []string {
	return a.attrs
}
