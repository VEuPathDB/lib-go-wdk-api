package step

import "github.com/VEuPathDB/lib-go-wdk-api/v0/optional"

type FullStep struct {
	ShortStep

	EstimatedSize optional.Int `json:"estimatedSize"`
}
