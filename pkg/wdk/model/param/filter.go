package param

import "lib-go-wdk-api/pkg/wdk/model/common"

type Filter struct {
	Base

	FilterDataTypeDisplayName string            `json:"filterDataTypeDisplayName"`
	Ontology                  []Ontology        `json:"ontology"`
	Values                    common.Properties `json:"values"`
	Type                      string            `json:"type"`
	HideEmptyOntologyNodes    bool              `json:"hideEmptyOntologyNodes"`
	SortLeavesBeforeBranches  bool              `json:"sortLeavesBeforeBranches"`
	CountPredictsAnswerCount  bool              `json:"countPredictsAnswerCount"`
	MinSelectedCount          int               `json:"minSelectedCount"`
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
