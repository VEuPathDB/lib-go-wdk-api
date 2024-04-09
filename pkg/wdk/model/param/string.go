package param

type String struct {
	Base

	Length      uint `json:"length"`
	IsMultiline bool `json:"isMultiline"`
	IsNumber    bool `json:"isNumber"`
}
