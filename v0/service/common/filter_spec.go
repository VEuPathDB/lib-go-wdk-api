package common

import "encoding/json"

type FilterSpec struct {
	Name     string          `json:"name"`
	Value    json.RawMessage `json:"value"`
	Disabled bool            `json:"disabled"`
}
