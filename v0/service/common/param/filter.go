package param

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
	. "github.com/VEuPathDB/lib-go-wdk-api/v0/service/common"
)

// Filter Param
//
// See org.gusdb.wdk.service.formatter.param.FilterParamNewFormatter
type Filter struct {
	Base

	FilterDataTypeDisplayName string     `json:"filterDataTypeDisplayName"`
	Ontology                  []Ontology `json:"ontology"`
	Values                    Properties `json:"values"`
	Type                      string     `json:"type"`
	HideEmptyOntologyNodes    bool       `json:"hideEmptyOntologyNodes"`
	SortLeavesBeforeBranches  bool       `json:"sortLeavesBeforeBranches"`
	CountPredictsAnswerCount  bool       `json:"countPredictsAnswerCount"`
	MinSelectedCount          int        `json:"minSelectedCount"`
}

type Ontology struct {
	Term        optional.String `json:"term"`
	Parent      optional.String `json:"parent"`
	Display     optional.String `json:"display"`
	Description optional.String `json:"description"`
	Type        optional.String `json:"type"`
	Units       optional.String `json:"units"`
	Precision   optional.Uint   `json:"precision"`
	IsRange     optional.Bool   `json:"isRange"`
}
