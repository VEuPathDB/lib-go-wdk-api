package optional

import "encoding/json"

func NewString(val string) String {
	return String{value: &val}
}

type String struct {
	value *string
}

func (s String) Exists() bool {
	return s.value != nil
}

func (s String) Get() string {
	return *s.value
}

func (s *String) Set(val string) {
	s.value = &val
}

func (s *String) Clear() {
	s.value = nil
}

func (s *String) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &s.value)
}

func (s String) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.value)
}
