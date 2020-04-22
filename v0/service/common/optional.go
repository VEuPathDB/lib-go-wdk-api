package common

import "encoding/json"

type OptionalString struct {
	value *string
}

func (o *OptionalString) Exists() bool {
	return o.value != nil
}

func (o *OptionalString) Get() string {
	return *o.value
}

func (o *OptionalString) Set(val string) {
	o.value = &val
}

func (o *OptionalString) Clear() {
	o.value = nil
}

func (o *OptionalString) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &o.value)
}

type OptionalUint struct {
	value *uint
}

func (o *OptionalUint) Exists() bool {
	return o.value != nil
}

func (o *OptionalUint) Get() uint {
	return *o.value
}

func (o *OptionalUint) Set(val uint) {
	o.value = &val
}

func (o *OptionalUint) Clear() {
	o.value = nil
}

func (o *OptionalUint) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &o.value)
}

type OptionalBool struct {
	value *bool
}

func (o *OptionalBool) Exists() bool {
	return o.value != nil
}

func (o *OptionalBool) Get() bool {
	return *o.value
}

func (o *OptionalBool) Set(val bool) {
	o.value = &val
}

func (o *OptionalBool) Clear() {
	o.value = nil
}

func (o *OptionalBool) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &o.value)
}
