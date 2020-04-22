package common

import (
	"encoding/json"
	"fmt"
)

// Align
type Align string

const (
	AlignCenter Align = "center"
	AlignLeft   Align = "left"
	AlignRight  Align = "right"
)

func (a *Align) UnmarshalJSON(bytes []byte) error {
	var tmp string

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	align := Align(tmp)

	switch align {
	case AlignCenter:
		fallthrough
	case AlignLeft:
		fallthrough
	case AlignRight:
		*a = align
	default:
		return fmt.Errorf("invalid align value \"%s\"", tmp)
	}

	return nil
}

type OptionalAlign struct {
	value *Align
}

func (o *OptionalAlign) Exists() bool {
	return o.value != nil
}

func (o *OptionalAlign) Get() Align {
	return *o.value
}

func (o *OptionalAlign) Set(val Align) {
	o.value = &val
}

func (o *OptionalAlign) Clear() {
	o.value = nil
}

func (o *OptionalAlign) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &o.value)
}
