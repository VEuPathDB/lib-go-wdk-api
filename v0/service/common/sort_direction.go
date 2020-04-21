package common

import (
	"encoding/json"
	"fmt"
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
		return fmt.Errorf("invalid sort direction \"%s\"", tmp)
	}

	return nil
}
