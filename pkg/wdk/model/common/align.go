package common

const (
	AlignCenter Align = "center"
	AlignLeft   Align = "left"
	AlignRight  Align = "right"
)

type Align string

func (a Align) Ptr() *Align {
	return &a
}
