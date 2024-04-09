package param

type Date struct {
	Base

	MinDate string `json:"minDate"`
	MaxDate string `json:"maxDate"`
}

type DateRange struct {
	Base

	MinDate string `json:"minDate"`
	MaxDate string `json:"maxDate"`
}
