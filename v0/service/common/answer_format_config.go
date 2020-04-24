package common

type AnswerFormatConfig struct {
	Pagination struct {
		Offset     uint `json:"offset"`
		NumRecords uint `json:"numRecords"`
	} `json:"pagination"`

	Attributes AllOrList     `json:"attributes"`
	Tables     AllOrList     `json:"tables"`
	Sorting    AttributeSort `json:"sorting"`
}
