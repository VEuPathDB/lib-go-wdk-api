package answer

import "github.com/VEuPathDB/lib-go-wdk-api/v0/service/common"

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

func (p *PostResponseMeta) HasRecordClassName() bool {
	return p.RecordClassName != nil
}

func (p *PostResponseMeta) GetRecordClassName() string {
	return *p.RecordClassName
}

func (p *PostResponseMeta) HasTotalCount() bool {
	return p.TotalCount != nil
}

func (p *PostResponseMeta) GetTotalCount() uint {
	return *p.TotalCount
}

func (p *PostResponseMeta) HasResponseCount() bool {
	return p.ResponseCount != nil
}

func (p *PostResponseMeta) GetResponseCount() uint {
	return *p.ResponseCount
}
