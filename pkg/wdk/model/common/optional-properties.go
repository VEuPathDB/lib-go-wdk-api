package common

func NewOptionalProperties(props Properties) OptionalProperties

type OptionalProperties struct {
	// contains filtered or unexported fields
}

func (o *OptionalProperties) Clear() {}

func (o OptionalProperties) Exists() bool {}

func (o OptionalProperties) Get() Properties {}

func (o OptionalProperties) MarshalJSON() ([]byte, error) {}

func (o *OptionalProperties) Set(val Properties) {}

func (o *OptionalProperties) UnmarshalJSON(bytes []byte) error {}
