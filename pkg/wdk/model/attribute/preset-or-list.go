package attribute

type PresetOrList struct {
	// contains filtered or unexported fields
}

func (a *PresetOrList) AllMarker() Preset {}

func (a *PresetOrList) IsAll() bool {}

func (a *PresetOrList) IsList() bool {}

func (a *PresetOrList) List() []string {}

func (a *PresetOrList) MarshalJSON() ([]byte, error) {}

func (a *PresetOrList) SetAll(mark Preset) {}

func (a *PresetOrList) SetList(list []string) {}

func (a *PresetOrList) UnmarshalJSON(bytes []byte) error {}
