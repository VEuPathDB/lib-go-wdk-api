package validation

import (
	"encoding/json"
	"fmt"
)

const (
	errBadLevel = `invalid validation level "%s"`
)

// Level represents the validation level under which the
// containing object was validated.
//
// See org.gusdb.fgputil.validation.ValidationLevel
type Level string

const (
	LevelNone        Level = "NONE"
	LevelSyntactic   Level = "SYNTACTIC"
	LevelDisplayable Level = "DISPLAYABLE"
	LevelSemantic    Level = "SEMANTIC"
	LevelRunnable    Level = "RUNNABLE"
)

var validLevels = []Level{
	LevelNone,
	LevelDisplayable,
	LevelSyntactic,
	LevelSemantic,
	LevelRunnable,
}

func (v Level) IsValid() bool {
	for i := range validLevels {
		if v == validLevels[i] {
			return true
		}
	}
	return false
}

func (v *Level) UnmarshalJSON(bytes []byte) error {
	var tmp string

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	vl := Level(tmp)

	if !vl.IsValid() {
		return fmt.Errorf(errBadLevel, tmp)
	}

	*v = vl
	return nil
}

func (v Level) MarshalJSON() ([]byte, error) {
	if !v.IsValid() {
		return nil, fmt.Errorf(errBadLevel, v)
	}
	return json.Marshal(string(v))
}
