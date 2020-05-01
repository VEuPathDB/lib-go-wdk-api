package report

import (
	"encoding/json"
	"fmt"
)

const (
	errBadAttType = "invalid fasta attachment type %s"
)

type FastaAttachmentType string

const (
	FastaAttachmentText FastaAttachmentType = "text"
)

func (f FastaAttachmentType) IsValid() bool {
	return f == FastaAttachmentText
}

func (f *FastaAttachmentType) UnmarshalJSON(bytes []byte) error {
	var tmp string
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}
	fat := FastaAttachmentType(tmp)
	if !fat.IsValid() {
		return fmt.Errorf(errBadAttType, tmp)
	}
	*f = FastaAttachmentType(tmp)
	return nil
}

func (f FastaAttachmentType) MarshalJSON() ([]byte, error) {
	if !f.IsValid() {
		return nil, fmt.Errorf(errBadAttType, f)
	}
	return json.Marshal(string(f))
}

type FastaReportConfig struct {
	AttachmentType *FastaAttachmentType `json:"attachmentType,omitempty"`
}

func (f *FastaReportConfig) ValidateReportConfig() error {
	if f.AttachmentType != nil && !f.AttachmentType.IsValid() {
		return fmt.Errorf(errBadAttType, *f.AttachmentType)
	}
	return nil
}
