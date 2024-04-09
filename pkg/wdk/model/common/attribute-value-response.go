package common

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
	// contains filtered or unexported fields
}

func (a *AttributeValueResponse) MarshalJSON() ([]byte, error)

func (a *AttributeValueResponse) UnmarshalJSON(bytes []byte) error
