package searches

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
	. "github.com/VEuPathDB/lib-go-wdk-api/v0/service/common"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/common/param"
)

//  "definitions": {
//    "properties": {
//      "type": "object",
//      "additionalProperties": {
//        "$ref": "../../../includes/string-array.json"
//      }
//    },
//    "filter-def-array": {
//      "type": "array",
//      "items": {
//        "type": "object",
//        "properties": {
//          "name": {
//            "type": "string"
//          },
//          "displayName": {
//            "type": "string"
//          },
//          "description": {
//            "type": "string"
//          },
//          "isViewOnly": {
//            "type": "boolean"
//          }
//        },
//        "required": [
//          "name",
//          "isViewOnly"
//        ]
//      }
//    }
//  },

//    "urlSegment",
//    "shortDisplayName",
//    "outputRecordClassName",
//    "defaultSorting",
//    "summaryViewPlugins",
//    "properties",
//    "paramNames"
type SearchResponse struct {
	/* Required Fields */

	DefaultAttributes       []string          `json:"defaultAttributes"`
	DefaultSummaryView      string            `json:"defaultSummaryView"`
	NoSummaryOnSingleRecord bool              `json:"noSummaryOnSingleRecord"`
	DisplayName             string            `json:"displayName"`
	FullName                string            `json:"fullName"`
	Filters                 []string          `json:"filters"`
	IsAnalyzable            bool              `json:"isAnalyzable"`
	DynamicAttributes       []RecordAttribute `json:"dynamicAttributes"`
	Groups                  []Group           `json:"groups"`

	/* Optional Fields */

	Description                           optional.String    `json:"description"`
	NewBuild                              optional.String    `json:"newBuild"`
	AllowedPrimaryInputRecordClassNames   []string           `json:"allowedPrimaryInputRecordClassNames"`
	AllowedSecondaryInputRecordClassNames []string           `json:"allowedSecondaryInputRecordClassNames"`
	Parameters                            []param.MultiParam `json:"parameters"`
	//    "parameters": {
	//          "type": "array",
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
	//    },
	//    "properties": {
	//      "$ref": "#/definitions/properties"
	//    },
	//    "outputRecordClassName": {
	//      "type": "string"
	//    },
	//    "shortDisplayName": {
	//      "type": "string"
	//    },
	//    "summary": {
	//      "type": "string"
	//    },
	//    "summaryViewPlugins": {
	//      "type": "array",
	//      "items": {
	//        "type": "object",
	//        "additionalProperties": false,
	//        "properties": {
	//          "description": {
	//            "type": "string"
	//          },
	//          "displayName": {
	//            "type": "string"
	//          },
	//          "name": {
	//            "type": "string"
	//          }
	//        }
	//      }
	//    },
	//    "urlSegment": {
	//      "type": "string"
	//    },
	//    "iconName": {
	//      "type": "string"
	//    },
	//    "help": {
	//      "type": "string"
	//    },
	//    "reviseBuild": {
	//      "type": "string"
	//    },
	//    "defaultSorting": {
	//      "$ref": "../../includes/sorting-spec.json"
	//    },
	//    "filters": {
	//      "$ref": "#/definitions/filter-def-array"
	//    },
	//    "paramNames": {
	//      "type": "array",
	//      "items": {
	//         "type": "string"
	//      }
	//    },
	//    "queryName": {
	//      "type": "string"
	//    }
}

type Group struct {
	/* Required Fields */

	Description string `json:"description"`
	DisplayName string `json:"displayName"`
	DisplayType string `json:"displayType"`
	Name        string `json:"name"`
	IsVisible   bool   `json:"isVisible"`

	/* Optional Fields */

	Parameters []string `json:"parameters"`
}
