package recordTypes

// Schema:
//
//     {
//       "$schema": "http://json-schema.org/draft-04/schema",
//       "oneOf": [
//         {
//           "$ref": "../../includes/string-array.json"
//         },
//         {
//           "type": "array",
//           "items": {
//             "$ref": "name/get.json"
//           }
//         }
//       ]
//     }


// NameList contains a list of record type names
type NameList []string

type ExpandedList []RecordTypeResponse