package param

type MultiParam struct {
	kind  Kind
	value interface{}
}

//          "items": [
//            {
//              "$ref": "../../includes/params/filter-param.json"
//            },
//            {
//              "$ref": "../../includes/params/answer-param.json"
//            },
//            {
//              "$ref": "../../includes/params/enum-param.json"
//            },
//            {
//              "$ref": "../../includes/params/string-param.json"
//            },
//            {
//              "$ref": "../../includes/params/number-param.json"
//            },
//            {
//              "$ref": "../../includes/params/date-param.json"
//            },
//            {
//              "$ref": "../../includes/params/dataset-param.json"
//            },
//            {
//              "$ref": "../../includes/params/filter-param.json"
//            },
//            {
//              "$ref": "../../includes/params/string-param.json"
//            },
//            {
//              "$ref": "../../includes/params/timestamp-param.json"
//            }
//          ]
