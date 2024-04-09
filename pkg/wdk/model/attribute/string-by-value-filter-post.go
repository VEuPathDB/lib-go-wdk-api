package attribute

type StringByValueFilterPost struct {
	// Pattern
	//
	// TODO: Explain this
	Pattern *string `json:"pattern,omitempty"`

	// Values is a list of values that the filtered rows must
	// match.
	Values []string `json:"values,omitempty"`
}
