package search

import "lib-go-wdk-api/pkg/wdk/model/validation"

type ValidatedSearch struct {
	SearchData FullSearch        `json:"searchData"`
	Validation validation.Bundle `json:"validation"`
}
