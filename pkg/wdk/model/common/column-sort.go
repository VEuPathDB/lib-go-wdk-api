package common

type ColumnSort struct {
	Name      string        `json:"name"`
	Direction SortDirection `json:"direction"`
}
