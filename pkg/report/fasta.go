package report

const (
	FastaAttachmentText FastaAttachmentType = "text"
)

type FastaAttachmentType string

func (f FastaAttachmentType) IsValid() bool

func (f FastaAttachmentType) MarshalJSON() ([]byte, error)

func (f *FastaAttachmentType) UnmarshalJSON(bytes []byte) error

type FastaReportConfig struct {
	AttachmentType *FastaAttachmentType `json:"attachmentType,omitempty"`
}

func (f *FastaReportConfig) ValidateReportConfig() error
