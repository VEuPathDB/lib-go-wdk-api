package wdk

type SearchApi interface {
	GetColumnNames() ([]string, error)
	MustGetColumnNames() []string

	GetColumns() ([]attribute.Field, error)
	MustGetColumns() []attribute.Field

	ColumnApiFor(columnName string) ColumnApi
}
