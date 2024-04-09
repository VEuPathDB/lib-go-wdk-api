package strategy

type CopyRequest struct {
	SourceStrategySignature string `json:"sourceStrategySignature"`
}

type CopyResponse struct {
	Id uint `json:"id"`
}
