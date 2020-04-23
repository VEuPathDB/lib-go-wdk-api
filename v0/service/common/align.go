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

func (a Align) IsValid() bool {
	return a == AlignCenter || a == AlignLeft || a == AlignRight
}

func (a *Align) UnmarshalJSON(bytes []byte) error {
	var tmp string

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	align := Align(tmp)

	if !align.IsValid() {
		return fmt.Errorf("invalid align value \"%s\"", tmp)
	}

	*a = align
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
