package validation

type Bundle struct {
	Level   Level         `json:"level"`
	IsValid bool          `json:"isValid"`
	Errors  *BundleErrors `json:"errors"`
}

func (v *Bundle) HasErrors() bool {
	return v.Errors != nil &&
		(len(v.Errors.General) > 0 || len(v.Errors.ByKey) > 0)
}

type BundleErrors struct {
	General []string            `json:"general"`
	ByKey   map[string][]string `json:"byKey"`
}
