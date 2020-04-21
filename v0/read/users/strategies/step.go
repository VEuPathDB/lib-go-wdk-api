package strategies

type Step struct {
	StepId         uint  `json:"stepId"`
	PrimaryInput   *Step `json:"primaryInput"`
	SecondaryInput *Step `json:"secondaryInput"`
}

func (s *Step) HasPrimaryInput() bool {
	return s.PrimaryInput != nil
}

func (s *Step) HasSecondaryInput() bool {
	return s.SecondaryInput != nil
}
