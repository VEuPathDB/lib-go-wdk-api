package answer

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/common"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
)

type PostResponse struct {
	Records []common.RecordInstanceResponse `json:"records"`
	Meta    PostResponseMeta                `json:"meta"`
}

type PostResponseMeta struct {
	RecordClassName        optional.String `json:"recordClassName"`
	TotalCount             optional.Uint   `json:"totalCount"`
	ResponseCount          optional.Uint   `json:"responseCount"`
	Tables                 []string        `json:"tables"`
	Attributes             []string        `json:"attributes"`
	CachePreviouslyExisted bool            `json:"cachePreviouslyExisted"`
}
