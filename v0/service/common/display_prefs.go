package common

type DisplayPrefs struct {
	ColumnSelection []string     `json:"columnSelection"`
	SortColumns     []ColumnSort `json:"sortColumns"`
}
