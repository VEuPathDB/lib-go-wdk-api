package optional

import "encoding/json"

func NewUint64(u uint64) Uint64 {
	return Uint64{&u}
}

type Uint64 struct {
	value *uint64
}

func (u Uint64) Exists() bool {
	return u.value != nil
}

func (u Uint64) Get() uint64 {
	return *u.value
}

func (u *Uint64) Set(val uint64) {
	u.value = &val
}

func (u *Uint64) Clear() {
	u.value = nil
}

func (u *Uint64) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &u.value)
}

func (u Uint64) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
}
