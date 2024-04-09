package common

type RecordInstanceResponse struct {
	Id              []StringKvPair           `json:"id"`
	DisplayName     string                   `json:"displayName"`
	RecordClassName string                   `json:"recordClassName"`
	Attributes      []AttributeValueResponse `json:"attributes"`
	Tables          RIRTables                `json:"tables"`
	TableErrors     []string                 `json:"tableErrors"`
}
