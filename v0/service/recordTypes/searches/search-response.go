package searches

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
	. "github.com/VEuPathDB/lib-go-wdk-api/v0/service/common"
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

type SearchResponse struct {
	/* Required Fields */

	DefaultAttributes       []string `json:"defaultAttributes"`
	DefaultSummaryView      string   `json:"defaultSummaryView"`
	NoSummaryOnSingleRecord bool     `json:"noSummaryOnSingleRecord"`
	DisplayName             string   `json:"displayName"`
	FullName                string   `json:"fullName"`

	// Filters has conflicting definitions, the other is:
	//    "filters": {
	//      "$ref": "#/definitions/filter-def-array"
	//    },
	Filters           []string          `json:"filters"`
	IsAnalyzable      bool              `json:"isAnalyzable"`
	DynamicAttributes []RecordAttribute `json:"dynamicAttributes"`
	Groups            []Group           `json:"groups"`
	//Parameters                            []param.MultiParam `json:"parameters"`
	Properties            Properties `json:"properties"`
	OutputRecordClassName string     `json:"outputRecordClassName"`
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

	//
	UrlSegment       string          `json:"urlSegment"`
	ShortDisplayName string          `json:"shortDisplayName"`
	DefaultSorting   []AttributeSort `json:"defaultSorting"`
	ParamNames       []string        `json:"paramNames"`

	/* Optional Fields */

	Description                           optional.String `json:"description"`
	NewBuild                              optional.String `json:"newBuild"`
	AllowedPrimaryInputRecordClassNames   []string        `json:"allowedPrimaryInputRecordClassNames"`
	AllowedSecondaryInputRecordClassNames []string        `json:"allowedSecondaryInputRecordClassNames"`
	Summary                               optional.String `json:"summary"`
	IconName                              optional.String `json:"iconName"`
	Help                                  optional.String `json:"help"`
	ReviseBuild                           optional.String `json:"reviseBuild"`
	Query                                 optional.String `json:"query"`
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
