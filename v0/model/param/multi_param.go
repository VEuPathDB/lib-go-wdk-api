package param

import (
	"encoding/json"
	"errors"
	"regexp"
)

const (
	errNoType = "no param type found"
)

var (
	typePattern = regexp.MustCompile(`"type": *"([^"])"`)
)

type MultiParam struct {
	kind  Kind
	value interface{}
}

func (m *MultiParam) UnmarshalJSON(bytes []byte) error {
	matches := typePattern.FindSubmatchIndex(bytes)
	if matches == nil {
		return errors.New(errNoType)
	}

	var kind Kind
	if err := kind.UnmarshalJSON(bytes[matches[2]:matches[3]]); err != nil {
		return err
	}

	m.kind = kind

	switch kind {
	case KindAnswer:
		var param Answer
		if err := json.Unmarshal(bytes, &param); err != nil {
			return err
		}
		m.value = param
	case KindFilter:
		var param Filter
		if err := json.Unmarshal(bytes, &param); err != nil {
			return err
		}
		m.value = param
	case KindDataset:
		var param Dataset
		if err := json.Unmarshal(bytes, &param); err != nil {
			return err
		}
		m.value = param
	case KindDate:
		var param Date
		if err := json.Unmarshal(bytes, &param); err != nil {
			return err
		}
		m.value = param
	case KindDateRange:
		var param DateRange
		if err := json.Unmarshal(bytes, &param); err != nil {
			return err
		}
		m.value = param
	case KindMultiVocab, KindSingleVocab:
		var param Enum
		if err := json.Unmarshal(bytes, &param); err != nil {
			return err
		}
		m.value = param
	case KindNumber:
		var param Number
		if err := json.Unmarshal(bytes, &param); err != nil {
			return err
		}
		m.value = param
	case KindNumberRange:
		var param NumberRange
		if err := json.Unmarshal(bytes, &param); err != nil {
			return err
		}
		m.value = param
	case KindString:
		var param String
		if err := json.Unmarshal(bytes, &param); err != nil {
			return err
		}
		m.value = param
	case KindTimestamp:
		var param Timestamp
		if err := json.Unmarshal(bytes, &param); err != nil {
			return err
		}
		m.value = param
	}
	return nil
}
