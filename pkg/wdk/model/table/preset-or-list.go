package table

type PresetOrList struct {
	// contains filtered or unexported fields
}

func (a *PresetOrList) IsAll() bool

func (a *PresetOrList) IsList() bool

func (a *PresetOrList) List() []string

func (a *PresetOrList) MarshalJSON() ([]byte, error)

func (a *PresetOrList) SetAll()

func (a *PresetOrList) SetList(list []string)

func (a *PresetOrList) UnmarshalJSON(bytes []byte) error
