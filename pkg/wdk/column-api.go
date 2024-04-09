package wdk

type ColumnApi interface {
	GetReports() ([]string, error)
	MustGetReports() []string

	GetFilters() ([]string, error)
	MustGetFilters() []string

	GetReportPostSchema(reportName string) ([]byte, error)
	MustGetReportPostSchema(reportName string) []byte

	GetFilterPostSchema(filterName string) ([]byte, error)
	MustGetFilterPostSchema(filterName string) []byte
}
