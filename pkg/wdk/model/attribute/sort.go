package attribute

import "lib-go-wdk-api/pkg/wdk/model/common"

type Sort struct {
	AttributeName string               `json:"attributeName"`
	Direction     common.SortDirection `json:"direction"`
}
