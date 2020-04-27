package attribute

import (
	"encoding/json"
	"fmt"
)

const (
	errBadFdt = `invalid field data type "%s"`
)

type FieldDataType string

const (
	FieldDataTypeString FieldDataType = "string"
	FieldDataTypeNumber FieldDataType = "number"
	FieldDataTypeDate   FieldDataType = "date"
	FieldDataTypeOther  FieldDataType = "other"
)

var validFieldDataTypes = []FieldDataType{
	FieldDataTypeString,
	FieldDataTypeNumber,
	FieldDataTypeDate,
	FieldDataTypeOther,
}

func (f FieldDataType) IsValid() bool {
	for i := range validFieldDataTypes {
		if f == validFieldDataTypes[i] {
			return true
		}
	}
	return false
}

func (f *FieldDataType) UnmarshalJSON(bytes []byte) error {
	var tmp string

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	scope := FieldDataType(tmp)

	if !scope.IsValid() {
		return fmt.Errorf(errBadFdt, tmp)
	}

	*f = scope
	return nil
}

func (f FieldDataType) MarshalJSON() ([]byte, error) {
	if !f.IsValid() {
		return nil, fmt.Errorf(errBadFdt, f)
	}
	return json.Marshal(string(f))
}
