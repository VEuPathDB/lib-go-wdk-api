package validation

import (
	"fmt"
	"strings"
)

const (
	LevelNone        Level = "NONE"
	LevelSyntactic   Level = "SYNTACTIC"
	LevelDisplayable Level = "DISPLAYABLE"
	LevelSemantic    Level = "SEMANTIC"
	LevelRunnable    Level = "RUNNABLE"
)

type Level string

func (v Level) IsValid() bool {
	switch v {
	case LevelNone, LevelSyntactic, LevelDisplayable, LevelSemantic, LevelRunnable:
		return true
	default:
		return false
	}
}

func (v Level) MarshalJSON() ([]byte, error) {
	out := make([]byte, len(v)+2)
	out[0], out[len(out)-1] = '"', '"'
	copy(out[1:], v)
	return out, nil
}

func (v *Level) UnmarshalJSON(bytes []byte) error {
	tmp := Level(strings.Trim(string(bytes), `"`))

	if !tmp.IsValid() {
		return fmt.Errorf("invalid validation level value: %s", tmp)
	}

	*v = tmp

	return nil
}
