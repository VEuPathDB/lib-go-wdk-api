package param

import (
	"encoding/json"
)

type MultiParam struct {
	Base

	value interface{}
}

type mpAlias MultiParam

func (m *MultiParam) UnmarshalJSON(bytes []byte) error {

	var kind Kind
	if err := json.Unmarshal(bytes, (*mpAlias)(m)); err != nil {
		return err
	}

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
