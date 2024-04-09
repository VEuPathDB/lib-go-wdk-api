package param

const (
	KindAnswer      Kind = "input-step"
	KindDate        Kind = "date"
	KindDateRange   Kind = "date-range"
	KindFilter      Kind = "filter"
	KindDataset     Kind = "input-dataset"
	KindNumber      Kind = "number"
	KindNumberRange Kind = "number-range"
	KindString      Kind = "string"
	KindTimestamp   Kind = "timestamp"
	KindSingleVocab Kind = "single-pick-vocabulary"
	KindMultiVocab  Kind = "multi-pick-vocabulary"
)

type Kind string
