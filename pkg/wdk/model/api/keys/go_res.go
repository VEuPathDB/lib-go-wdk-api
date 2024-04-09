package keys

// Schema:
//   {
//     "$schema": "http://json-schema.org/draft-04/schema#",
//     "type": "object",
//     "patternProperties": {
//       "^\\w+$": {
//         "type": "string"
//       }
//     },
//     "additionalProperties": false
//   }
type GetResponse map[string]string
