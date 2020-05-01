package param

import (
	. "github.com/VEuPathDB/lib-go-wdk-api/v0/model/common"
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
	Term        *string `json:"term"`
	Parent      *string `json:"parent"`
	Display     *string `json:"display"`
	Description *string `json:"description"`
	Type        *string `json:"type"`
	Units       *string `json:"units"`
	Precision   *uint   `json:"precision"`
	IsRange     *bool   `json:"isRange"`
}
