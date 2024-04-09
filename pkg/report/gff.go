package report

type GenBankReportConfig struct{}

func (g *GenBankReportConfig) ValidateReportConfig() error

// Gff3CachedReportConfig
//
// See org.apidb.apicommon.model.report.Gff3CachedReporter
type Gff3CachedReportConfig struct {
	AttachmentType *string `json:"downloadType,omitempty"`
	HasTranscript  *bool   `json:"hasTranscript,omitempty"`
	HasProtein     *bool   `json:"hasProtein,omitempty"`
}

func (g *Gff3CachedReportConfig) ValidateReportConfig() error

type Gff3ReportConfig struct {
	AttachmentType *string `json:"downloadType,omitempty"`
	HasTranscript  *bool   `json:"hasTranscript,omitempty"`
	HasProtein     *bool   `json:"hasProtein,omitempty"`
}

func (g *Gff3ReportConfig) ValidateReportConfig() error
