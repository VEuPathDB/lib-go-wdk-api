package common

import (
	"encoding/json"
	"fmt"
)

type ValidationLevel string

const (
	ValidationLevelNone        ValidationLevel = "NONE"
	ValidationLevelUnspecified ValidationLevel = "UNSPECIFIED"
	ValidationLevelSyntactic   ValidationLevel = "SYNTACTIC"
	ValidationLevelSemantic    ValidationLevel = "SEMANTIC"
	ValidationLevelRunnable    ValidationLevel = "RUNNABLE"
)

func (v *ValidationLevel) UnmarshalJSON(bytes []byte) error {
	var tmp string
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}
	vl := ValidationLevel(tmp)
	switch vl {
	case ValidationLevelNone, ValidationLevelUnspecified,
		ValidationLevelSyntactic, ValidationLevelSemantic, ValidationLevelRunnable:
		*v = vl
	default:
		return fmt.Errorf("invalid validation level \"%s\"", tmp)
	}
	return nil
}

type ValidationBundle struct {
	Level   ValidationLevel `json:"level"`
	IsValid bool            `json:"isValid"`
	Errors  *struct{
		General []string `json:"general"`
		ByKey   map[string][]string `json:"byKey"`
	} `json:"errors"`
}

func (v *ValidationBundle) HasErrors() bool {
	return v.Errors != nil
}
