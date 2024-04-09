package param

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"lib-go-wdk-api/pkg/wdk/model/common"
)

var mpKindRegex = regexp.MustCompile(`"type"[^:]*:[^"]"([\w\-]+)"`)

type MultiParam struct {
	base  *Base
	value any
}

func (m MultiParam) Name() string { return m.base.Name }

func (m MultiParam) DisplayName() string { return m.base.DisplayName }

func (m MultiParam) Help() string { return m.base.Help }

func (m MultiParam) Kind() Kind { return m.base.Kind }

func (m MultiParam) IsVisible() bool { return m.base.IsVisible }

func (m MultiParam) Group() string { return m.base.Group }

func (m MultiParam) IsReadOnly() bool { return m.base.IsReadOnly }

func (m MultiParam) AllowEmptyValue() bool { return m.base.AllowEmptyValue }

func (m MultiParam) VisibleHelp() string { return m.base.VisibleHelp }

func (m MultiParam) DependentParams() []string { return m.base.DependentParams }

func (m MultiParam) InitialDisplayValue() *string { return m.base.InitialDisplayValue }

func (m MultiParam) PropertyLists() common.Properties { return m.base.PropertyLists }

func (m *MultiParam) AsAnswerParam() Answer { return m.value.(Answer) }

func (m *MultiParam) AsKindDataset() Dataset { return m.value.(Dataset) }

func (m *MultiParam) AsKindDate() Date { return m.value.(Date) }

func (m *MultiParam) AsKindDateRange() DateRange { return m.value.(DateRange) }

func (m *MultiParam) AsKindEnum() Enum { return m.value.(Enum) }

func (m *MultiParam) AsKindFilter() Filter { return m.value.(Filter) }

func (m *MultiParam) AsKindMultiVocab() Enum { return m.value.(Enum) }

func (m *MultiParam) AsKindNumber() Number { return m.value.(Number) }

func (m *MultiParam) AsKindNumberRange() NumberRange { return m.value.(NumberRange) }

func (m *MultiParam) AsKindSingleVocab() Enum { return m.value.(Enum) }

func (m *MultiParam) AsKindTimestamp() Timestamp { return m.value.(Timestamp) }

func (m *MultiParam) RawParam() any { return m.value }

func (m *MultiParam) UnmarshalJSON(bytes []byte) error {
	res := mpKindRegex.FindSubmatch(bytes)

	if len(res) == 0 {
		return errors.New(`param JSON did not contain a "type" field`)
	}

	switch Kind(res[1]) {

	case KindAnswer:
		var tmp Answer
		if err := json.Unmarshal(bytes, &tmp); err != nil {
			return err
		}
		m.base = &tmp.Base
		m.value = tmp

	case KindDataset:
		var tmp Dataset
		if err := json.Unmarshal(bytes, &tmp); err != nil {
			return err
		}
		m.base = &tmp.Base
		m.value = tmp

	case KindDate:
		var tmp Date
		if err := json.Unmarshal(bytes, &tmp); err != nil {
			return err
		}
		m.base = &tmp.Base
		m.value = tmp

	case KindDateRange:
		var tmp DateRange
		if err := json.Unmarshal(bytes, &tmp); err != nil {
			return err
		}
		m.base = &tmp.Base
		m.value = tmp

	case KindFilter:
		var tmp Filter
		if err := json.Unmarshal(bytes, &tmp); err != nil {
			return err
		}
		m.base = &tmp.Base
		m.value = tmp

	case KindNumber:
		var tmp Number
		if err := json.Unmarshal(bytes, &tmp); err != nil {
			return err
		}
		m.base = &tmp.Base
		m.value = tmp

	case KindNumberRange:
		var tmp NumberRange
		if err := json.Unmarshal(bytes, &tmp); err != nil {
			return err
		}
		m.base = &tmp.Base
		m.value = tmp

	case KindString:
		var tmp String
		if err := json.Unmarshal(bytes, &tmp); err != nil {
			return err
		}
		m.base = &tmp.Base
		m.value = tmp

	case KindTimestamp:
		var tmp Timestamp
		if err := json.Unmarshal(bytes, &tmp); err != nil {
			return err
		}
		m.base = &tmp.Base
		m.value = tmp

	case KindMultiVocab, KindSingleVocab:
		var tmp Enum
		if err := json.Unmarshal(bytes, &tmp); err != nil {
			return err
		}
		m.base = &tmp.Base
		m.value = tmp

	default:
		return fmt.Errorf("unrecognized param type: %s", string(res[1]))
	}

	return nil
}
