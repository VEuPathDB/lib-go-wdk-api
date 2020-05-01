package attribute

import (
	"errors"
)

const (
	errSbvfpXor = "type StringByValueFilterPost must contain exactly one " +
		"of: a non-empty Pattern value or a non-empty Values slice"
	errNbvfpXor = "type NumberByValueFilterPost must contain exactly one " +
		"of: a non-empty Range value or a non-empty Values slice"
)

// StringByValueFilterPost defines the post body for running
// a "byValue" column/attribute filter.
//
// Value must be exactly one of the following properties, if
// neither or both are set, validation will fail.
type StringByValueFilterPost struct {
	// Pattern
	//
	// TODO: Explain this
	Pattern *string `json:"pattern,omitempty"`

	// Values is a list of values that the filtered rows must
	// match.
	Values []string `json:"values,omitempty"`
}

// Validate verifies that exactly 1 property on
// StringByValueFilterPost exists and is non-empty.
func (s *StringByValueFilterPost) Validate() error {
	vals := len(s.Values) > 0
	patt := s.Pattern != nil && len(*s.Pattern) > 0

	if vals == patt {
		return errors.New(errSbvfpXor)
	}

	return nil
}

type NumberByValueFilterPost struct {
	Range  *NumberRange `json:"range,omitempty"`
	Values []float64    `json:"values,omitempty"`
}

func (n *NumberByValueFilterPost) Validate() error {
	ran := n.Range != nil
	val := len(n.Values) > 0

	if ran == val {
		return errors.New(errNbvfpXor)
	}

	return nil
}

type NumberRange struct {
	Min NumberRangeEnd `json:"min"`
	Max NumberRangeEnd `json:"max"`
}

type NumberRangeEnd struct {
	Value       float64 `json:"value"`
	IsInclusive bool    `json:"isInclusive"`
}
