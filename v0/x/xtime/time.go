package xtime

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	errBadTimeFormat = `datetime string "%s" does not match an expected format`
)

var timeFormats = []string{
	time.RFC3339Nano,
	"2006-01-02T15:04:05.999999999",
}

// Time is a custom alias for the standard library time type
// allowing the JSON unmarshaling to handle more than one
// time format.
type Time time.Time

func (t *Time) UnmarshalJSON(bytes []byte) error {
	var tmp string
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	for _, format := range timeFormats {
		tim, err := time.Parse(format, tmp)
		if err != nil {
			continue
		}
		*t = Time(tim)
		return nil
	}

	return fmt.Errorf(errBadTimeFormat, tmp)
}

