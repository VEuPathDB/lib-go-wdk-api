package xtime

import (
	"strings"
	"time"
)

const ISOLocalDateTime = "2006-01-02T15:04:05"

type DateTime time.Time

func (t *DateTime) UnmarshalJSON(bytes []byte) error {
	if parsed, err := time.Parse(ISOLocalDateTime, strings.Trim(string(bytes), `"`)); err != nil {
		return err
	} else {
		*t = DateTime(parsed)
	}

	return nil
}

func (t DateTime) AsGoTime() time.Time {
	return time.Time(t)
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(t).Format(ISOLocalDateTime)), nil
}
