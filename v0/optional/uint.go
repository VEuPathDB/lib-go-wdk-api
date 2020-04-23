package optional

import "encoding/json"

func NewUint(u uint) Uint {
	return Uint{&u}
}

type Uint struct {
	value *uint
}

func (u Uint) Exists() bool {
	return u.value != nil
}

func (u Uint) Get() uint {
	return *u.value
}

func (u *Uint) Set(val uint) {
	u.value = &val
}

func (u *Uint) Clear() {
	u.value = nil
}

func (u *Uint) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &u.value)
}

func (u Uint) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
}
