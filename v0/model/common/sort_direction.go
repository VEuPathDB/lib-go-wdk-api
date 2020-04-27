package common

import (
	"encoding/json"
	"fmt"
)

const (
	errBadSortDir = "invalid sort direction \"%s\""
)

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
		return fmt.Errorf(errBadSortDir, tmp)
	}

	return nil
}

func (s SortDirection) MarshalJSON() ([]byte, error) {
	if s != SortDirectionAsc && s != SortDirectionDesc {
		return nil, fmt.Errorf(errBadSortDir, s)
	}
	return json.Marshal(string(s))
}
