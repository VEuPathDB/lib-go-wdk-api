package common

import "encoding/json"

type ColumnFilters map[string]map[string]json.RawMessage

func (c ColumnFilters) Append(a, b string, val interface{}) error {
	bytes, err := json.Marshal(val)
	if err != nil {
		return err
	}

	if _, ok := c[a]; !ok {
		c[a] = make(map[string]json.RawMessage)
	}

	c[a][b] = bytes
	return nil
}
