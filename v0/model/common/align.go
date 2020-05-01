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

func (a Align) Ptr() *Align {
	return &a
}

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
