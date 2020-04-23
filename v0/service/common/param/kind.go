package param

import (
	"encoding/json"
	"fmt"
)

type Kind string

const (
	KindDate        Kind = "date"
	KindDateRange   Kind = "date-range"
	KindFilter      Kind = "filter"
	KindDataset     Kind = "input-dataset"
	KindAnswer      Kind = "input-step"
	KindNumber      Kind = "number"
	KindNumberRange Kind = "number-range"
	KindString      Kind = "string"
	KindTimestamp   Kind = "timestamp"
	KindSingleVocab Kind = "single-pick-vocabulary"
	KindMultiVocab  Kind = "multi-pick-vocabulary"
)

var validKinds = map[Kind]bool{
	KindDate:        true,
	KindDateRange:   true,
	KindFilter:      true,
	KindDataset:     true,
	KindAnswer:      true,
	KindNumber:      true,
	KindNumberRange: true,
	KindString:      true,
	KindTimestamp:   true,
	KindSingleVocab: true,
	KindMultiVocab:  true,
}

func (k *Kind) UnmarshalJSON(bytes []byte) error {
	var tmp string

	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	kind := Kind(tmp)
	if !kind.IsValid() {
		return fmt.Errorf("unrecognized param type \"%s\"", tmp)
	}

	*k = kind
	return nil
}

func (k Kind) MarshalJSON() ([]byte, error) {
	if !k.IsValid() {
		return nil, fmt.Errorf("unrecognized param type \"%s\"", k)
	}

	return json.Marshal(string(k))
}

func (k Kind) IsValid() bool {
	return validKinds[k]
}
