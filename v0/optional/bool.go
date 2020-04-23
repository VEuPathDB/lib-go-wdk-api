package optional

import "encoding/json"

func NewBool(b bool) Bool {
	return Bool{&b}
}

type Bool struct {
	value *bool
}

func (b Bool) Exists() bool {
	return b.value != nil
}

func (b Bool) Get() bool {
	return *b.value
}

func (b *Bool) Set(val bool) {
	b.value = &val
}

func (b *Bool) Clear() {
	b.value = nil
}

func (b *Bool) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &b.value)
}

func (b Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.value)
}
