package param

import (
	"encoding/json"
)

type mpAlias MultiParam

type MultiParam struct {
	Base

	value interface{}
}

func (m *MultiParam) RawParam() interface{} {
	return m.value
}

func (m *MultiParam) AsAnswerParam() Answer          { return m.value.(Answer) }
func (m *MultiParam) AsKindFilter() Filter           { return m.value.(Filter) }
func (m *MultiParam) AsKindDataset() Dataset         { return m.value.(Dataset) }
func (m *MultiParam) AsKindDate() Date               { return m.value.(Date) }
func (m *MultiParam) AsKindDateRange() DateRange     { return m.value.(DateRange) }
func (m *MultiParam) AsKindMultiVocab() Enum         { return m.value.(Enum) }
func (m *MultiParam) AsKindSingleVocab() Enum        { return m.value.(Enum) }
func (m *MultiParam) AsKindEnum() Enum               { return m.value.(Enum) }
func (m *MultiParam) AsKindNumber() Number           { return m.value.(Number) }
func (m *MultiParam) AsKindNumberRange() NumberRange { return m.value.(NumberRange) }
func (m *MultiParam) AsKindTimestamp() Timestamp     { return m.value.(Timestamp) }

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
