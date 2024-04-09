package answer

import "lib-go-wdk-api/pkg/wdk/model/common"

type Format struct {
	Name        string         `json:"name"`
	Type        string         `json:"type"`
	DisplayName string         `json:"displayName"`
	Description *string        `json:"description"`
	IsInReport  bool           `json:"isInReport"`
	Scopes      []common.Scope `json:"scopes"`
}
