package attribute

import "github.com/VEuPathDB/lib-go-wdk-api/v0/model/common"

type Sort struct {
	AttributeName string               `json:"attributeName"`
	Direction     common.SortDirection `json:"direction"`
}
