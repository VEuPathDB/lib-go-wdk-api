package report

// JsonReportConfig
//
// See org.gusdb.wdk.model.report.reporter.JSONReporter
type JsonReportConfig struct{ StandardReportConfig }

// StandardReportConfig
//
// See org.gusdb.wdk.model.report.config.StandardConfig#configure(org.json.JSONObject)
type StandardReportConfig struct {
	IncludeEmptyTables  *bool           `json:"includeEmptyTables,omitempty"`
	SelectAllFields     *bool           `json:"all-fields,omitempty"`
	SelectedFields      []string        `json:"selectedFields,omitempty"`
	SelectAllAttributes *bool           `json:"allAttributes,omitempty"`
	SelectedAttributes  []string        `json:"attributes,omitempty"`
	SelectAllTables     *bool           `json:"allTables,omitempty"`
	SelectedTables      []string        `json:"tables,omitempty"`
	AttachmentType      *string         `json:"downloadType,omitempty"`
	StreamStrategy      *StreamStrategy `json:"streamStrategy,omitempty"`
}

func (s *StandardReportConfig) ValidateReportConfig() error

const (
	StreamStratPaged StreamStrategy = "PAGED_ANSWER"
	StreamStratFile  StreamStrategy = "FILE_BASED"
	StreamStratSmart StreamStrategy = "SMART"
)

type StreamStrategy string

// XmlReportConfig
//
// See org.gusdb.wdk.model.report.reporter.XMLReporter
type XmlReportConfig struct{ StandardReportConfig }
