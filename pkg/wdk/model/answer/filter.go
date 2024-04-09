package answer

import "encoding/json"

type Filter struct {
	Name     string          `json:"name"`
	Value    json.RawMessage `json:"value"`
	Disabled bool            `json:"disabled"`
}
