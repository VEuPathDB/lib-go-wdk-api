package attribute

type NumberByValueFilterPost struct {
	Range  *NumberRange `json:"range,omitempty"`
	Values []float64    `json:"values,omitempty"`
}
