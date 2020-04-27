package common

import "encoding/json"

type Properties map[string][]string

func NewOptionalProperties(props Properties) OptionalProperties {
	return OptionalProperties{props}
}

type OptionalProperties struct {
	value Properties
}

func (o OptionalProperties) Exists() bool {
	return o.value != nil
}

func (o OptionalProperties) Get() Properties {
	return o.value
}

func (o *OptionalProperties) Set(val Properties) {
	o.value = val
}

func (o *OptionalProperties) Clear() {
	o.value = nil
}

func (o *OptionalProperties) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &o.value)
}

func (o OptionalProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}
