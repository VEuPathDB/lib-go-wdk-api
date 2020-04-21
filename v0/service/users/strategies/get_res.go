package strategies

import "github.com/VEuPathDB/lib-go-wdk-api/v0/service/common"

// Schema:
//
//     {
//       "$schema": "http://json-schema.org/draft-04/schema",
//       "type": "array",
//       "items": {
//         "$ref": "../../includes/strategy-no-step-tree.json",
//         "additionalProperties": false
//       }
//     }
type ListResponse []common.StrategyNoStepTree
