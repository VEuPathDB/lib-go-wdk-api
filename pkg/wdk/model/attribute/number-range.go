package attribute

type NumberRange struct {
	Min NumberRangeEnd `json:"min"`
	Max NumberRangeEnd `json:"max"`
}
