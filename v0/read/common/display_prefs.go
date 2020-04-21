package common

import (
	"encoding/json"
	"fmt"
)

type ColumnSort struct {
	Name      string        `json:"name"`
	Direction SortDirection `json:"direction"`
}

type SortDirection string

const (
	SortDirectionAsc  SortDirection = "ASC"
	SortDirectionDesc SortDirection = "DESC"
)

func (s *SortDirection) UnmarshalJSON(bytes []byte) error {
	var tmp string
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}
	sd := SortDirection(tmp)
	switch sd {
	case SortDirectionAsc, SortDirectionDesc:
		*s = sd
	default:
		return fmt.Errorf("invalid sort direction \"%s\"", tmp)
	}
	return nil
}

type DisplayPrefs struct {
	ColumnSelection []string `json:"columnSelection"`
	SortColumns     []ColumnSort `json:"sortColumns"`
}
