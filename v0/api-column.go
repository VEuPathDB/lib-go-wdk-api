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

type columnApi struct {
	report string
	search string
	column string
}

func (c *columnApi) GetReports() ([]string, error) {
	panic("implement me")
}

func (c *columnApi) MustGetReports() []string {
	panic("implement me")
}

func (c *columnApi) GetFilters() ([]string, error) {
	panic("implement me")
}

func (c *columnApi) MustGetFilters() []string {
	panic("implement me")
}

func (c *columnApi) GetReportPostSchema(name string) ([]byte, error) {
	panic("implement me")
}

func (c *columnApi) MustGetReportPostSchema(name string) []byte {
	panic("implement me")
}

func (c *columnApi) GetFilterPostSchema(name string) ([]byte, error) {
	panic("implement me")
}

func (c *columnApi) MustGetFilterPostSchema(name string) []byte {
	panic("implement me")
}
