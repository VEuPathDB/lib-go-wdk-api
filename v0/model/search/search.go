package search

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/param"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/validation"
)

type FullSearch struct {
	Base
	param.Container
}

type ValidatedSearch struct {
	SearchData FullSearch        `json:"searchData"`
	Validation validation.Bundle `json:"validation"`
}
