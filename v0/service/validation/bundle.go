package validation

type Bundle struct {
	Level   Level         `json:"level"`
	IsValid bool          `json:"isValid"`
	Errors  *BundleErrors `json:"errors"`
}

func (v *Bundle) HasErrors() bool {
	return v.Errors != nil
}

type BundleErrors struct {
	General []string            `json:"general"`
	ByKey   map[string][]string `json:"byKey"`
}
