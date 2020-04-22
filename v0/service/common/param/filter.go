package param

import . "github.com/VEuPathDB/lib-go-wdk-api/v0/service/common"

type Filter struct {
	Base

	Type                      string     `json:"type"`
	CountOnlyLeaves           bool       `json:"countOnlyLeaves"`
	FilterDataTypeDisplayName string     `json:"filterDataTypeDisplayName"`
	HideEmptyOntologyNodes    bool       `json:"hideEmptyOntologyNodes"`
	MinSelectedCount          bool       `json:"minSelectedCount"`
	Ontology                  []Ontology `json:"ontology"`
	Values                    Properties `json:"values"`
}

type Ontology struct {
	Term        OptionalString `json:"term"`
	Parent      OptionalString `json:"parent"`
	Display     OptionalString `json:"display"`
	Description OptionalString `json:"description"`
	Type        OptionalString `json:"type"`
	Units       OptionalString `json:"units"`
	IsRange     OptionalBool   `json:"isRange"`
	Precision   OptionalUint   `json:"precision"`
}
