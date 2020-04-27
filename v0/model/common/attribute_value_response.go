package common

import "encoding/json"

// Schema:
//   "oneOf": [
//     {
//       "type": "null"
//     },
//     {
//       "type": "object",
//       "additionalProperties": false,
//       "properties": {
//         "url": {
//           "type": "string"
//         },
//         "displayName": {
//           "type": "string"
//         }
//       },
//       "required": [
//         "url",
//         "displayName"
//       ]
//     },
//     {
//       "type": "object"
//     }
//   ]
type AttributeValueResponse struct {
	none   bool
	strong struct {
		Url         string `json:"url"`
		DisplayName string `json:"displayName"`
	}
	weak json.RawMessage
}

func (a *AttributeValueResponse) UnmarshalJSON(bytes []byte) error {
	if len(bytes) == 4 && string(bytes) == "null" {
		a.none = true
		return nil
	}

	var tmp map[string]string
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		a.weak = bytes
		return nil
	}

	if len(tmp) != 2 {
		a.weak = bytes
		return nil
	}

	url, ok1 := tmp["url"]
	dn, ok2 := tmp["displayName"]

	if !ok1 || !ok2 {
		a.weak = bytes
		return nil
	}

	a.strong.Url = url
	a.strong.DisplayName = dn
	return nil
}

func (a *AttributeValueResponse) MarshalJSON() ([]byte, error) {
	if a.none {
		return []byte{'n', 'u', 'l', 'l'}, nil
	}
	if a.weak != nil {
		return a.weak, nil
	}
	return json.Marshal(a.strong)
}
