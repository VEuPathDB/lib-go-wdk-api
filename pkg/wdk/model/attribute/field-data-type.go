package attribute

type FieldDataType string

const (
	FieldDataTypeString FieldDataType = "STRING"
	FieldDataTypeNumber FieldDataType = "NUMBER"
	FieldDataTypeDate   FieldDataType = "DATE"
	FieldDataTypeOther  FieldDataType = "OTHER"
)
