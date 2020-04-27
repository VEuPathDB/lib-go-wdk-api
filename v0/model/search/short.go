package search

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/param"
)

// BasicSearch
//
// See org.gusdb.wdk.service.formatter.QuestionFormatter#getQuestionJsonWithoutParams
type ShortSearch struct {
	Base
	param.BasicContainer
}
