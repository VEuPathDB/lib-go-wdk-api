package report

// AttributesTabularReportConfig
//
// See org.gusdb.wdk.model.report.reporter.AttributesTabularReporter
type AttributesTabularReportConfig struct{ TabularReportConfig }

// TableTabularReportConfig
//
// See org.gusdb.wdk.model.report.reporter.TableTabularReporter
type TableTabularReportConfig struct{ TabularReportConfig }

type TabularReportConfig struct {
	IncludeHeader *bool `json:"includeHeader,omitempty"`
}

func (t *TabularReportConfig) ValidateReportConfig() error

// TranscriptAttributesReportConfig
//
// See org.apidb.apicommon.model.report.TranscriptAttributesReporter
type TranscriptAttributesReportConfig struct{ AttributesTabularReportConfig }

// TranscriptTableReportConfig
//
// See org.apidb.apicommon.model.report.TranscriptTableReporter
type TranscriptTableReportConfig struct{ TableTabularReportConfig }
