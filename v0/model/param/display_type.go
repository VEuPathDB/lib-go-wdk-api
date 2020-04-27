package param

import (
	"encoding/json"
	"fmt"
)

const (
	errBadDisplayType = `unrecognized DisplayType value "%s"`
)

// DisplayType
//
// See org.gusdb.wdk.model.query.param.AbstractEnumParam.DisplayType
type DisplayType string

const (
	DisplayTypeSelect    DisplayType = "select"
	DisplayTypeCheckBox  DisplayType = "checkBox"
	DisplayTypeTreeBox   DisplayType = "treeBox"
	DisplayTypeTypeAhead DisplayType = "typeAhead"
)

func DisplayTypeValues() []DisplayType {
	return []DisplayType{DisplayTypeSelect, DisplayTypeCheckBox,
		DisplayTypeTreeBox, DisplayTypeTypeAhead}
}

func (d DisplayType) IsValid() bool {
	tmp := DisplayTypeValues()
	for i := range tmp {
		if d == tmp[i] {
			return true
		}
	}
	return false
}

func (d *DisplayType) UnmarshalJSON(bytes []byte) error {
	var tmp string

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	dt := DisplayType(tmp)

	if !dt.IsValid() {
		return fmt.Errorf(errBadDisplayType, tmp)
	}

	*d = dt
	return nil
}

func (d DisplayType) MarshalJSON() ([]byte, error) {
	if !d.IsValid() {
		return nil, fmt.Errorf(errBadDisplayType, d)
	}

	return json.Marshal(string(d))
}
