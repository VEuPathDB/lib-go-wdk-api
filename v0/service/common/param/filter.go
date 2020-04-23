package param

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
	. "github.com/VEuPathDB/lib-go-wdk-api/v0/service/common"
)

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
	Term        optional.String `json:"term"`
	Parent      optional.String `json:"parent"`
	Display     optional.String `json:"display"`
	Description optional.String `json:"description"`
	Type        optional.String `json:"type"`
	Units       optional.String `json:"units"`
	IsRange     optional.Bool   `json:"isRange"`
	Precision   optional.Uint   `json:"precision"`
}
