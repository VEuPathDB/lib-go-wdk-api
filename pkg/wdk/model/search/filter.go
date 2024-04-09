package search

type Filter struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	IsViewOnly  bool   `json:"isViewOnly"`
}
