package common

type ColumnSort struct {
	Name      string        `json:"name"`
	Direction SortDirection `json:"direction"`
}

type AttributeSort struct {
	AttributeName string        `json:"attributeName"`
	Direction     SortDirection `json:"direction"`
}
