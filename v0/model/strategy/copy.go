package strategy

type CopyResponse struct {
	Id uint `json:"id"`
}

type CopyRequest struct {
	SourceStrategySignature string `json:"sourceStrategySignature"`
}
