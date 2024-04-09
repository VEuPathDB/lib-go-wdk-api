package param

type Dataset struct {
	Base

	DefaultIdList   string          `json:"defaultIdList"`
	RecordClassName string          `json:"recordClassName"`
	Parsers         []DatasetParser `json:"parsers"`
}

type DatasetParser struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
}
