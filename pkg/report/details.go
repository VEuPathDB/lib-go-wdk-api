package report

import (
	"lib-go-wdk-api/pkg/wdk/model/attribute"
	"lib-go-wdk-api/pkg/wdk/model/table"
)

type DetailsReportConfig struct {
	Pagination         *Pagination             `json:"pagination,omitempty"`
	Attributes         *attribute.PresetOrList `json:"attributes,omitempty"`
	Tables             *table.PresetOrList     `json:"tables,omitempty"`
	Sorting            []attribute.Sort        `json:"sorting,omitempty"`
	ContentDisposition *string                 `json:"contentDisposition,omitempty"`
}

func (d *DetailsReportConfig) ValidateReportConfig() error

type Pagination struct {
	Offset uint `json:"offset"`

	// Count defines the max number of records to return in
	// the response.
	//
	// If count is less than 0, then all records will be returned.
	Count int `json:"numRecords"`
}
