package step

type FullStep struct {
	ShortStep

	EstimatedSize *int `json:"estimatedSize"`
}
