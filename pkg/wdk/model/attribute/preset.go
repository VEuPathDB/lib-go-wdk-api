package attribute

const (
	PresetAllAttributes     Preset = "__ALL_ATTRIBUTES__"
	PresetDefaultAttributes Preset = "__DEFAULT_ATTRIBUTES__"
)

type Preset string

func (a Preset) Ptr() *Preset {
	return &a
}
