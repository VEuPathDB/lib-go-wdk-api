package optional

import "encoding/json"

func NewInt(u int) Int {
	return Int{&u}
}

type Int struct {
	value *int
}

func (i Int) Exists() bool {
	return i.value != nil
}

func (i Int) Get() int {
	return *i.value
}

func (i *Int) Set(val int) {
	i.value = &val
}

func (i *Int) Clear() {
	i.value = nil
}

func (i *Int) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &i.value)
}

func (i Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.value)
}
