package answer

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/common"
)

type PostResponse struct {
	Records []common.RecordInstanceResponse `json:"records"`
	Meta    PostResponseMeta                `json:"meta"`
}

type PostResponseMeta struct {
	RecordClassName        *string  `json:"recordClassName"`
	TotalCount             *uint    `json:"totalCount"`
	ResponseCount          *uint    `json:"responseCount"`
	Tables                 []string `json:"tables"`
	Attributes             []string `json:"attributes"`
	CachePreviouslyExisted bool     `json:"cachePreviouslyExisted"`
}
