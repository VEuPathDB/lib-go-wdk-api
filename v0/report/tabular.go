package report

// TabularReportConfig
//
// See org.gusdb.wdk.model.report.reporter.AbstractTabularReporter
type TabularReportConfig struct {
	IncludeHeader *bool `json:"includeHeader,omitempty"`
}

func (t *TabularReportConfig) ValidateReportConfig() error {
	return nil
}

// AttributesTabularReportConfig
//
// See org.gusdb.wdk.model.report.reporter.AttributesTabularReporter
type AttributesTabularReportConfig struct{ TabularReportConfig }

// TableTabularReportConfig
//
// See org.gusdb.wdk.model.report.reporter.TableTabularReporter
type TableTabularReportConfig struct{ TabularReportConfig }

// TranscriptAttributesReportConfig
//
// See org.apidb.apicommon.model.report.TranscriptAttributesReporter
type TranscriptAttributesReportConfig struct{ AttributesTabularReportConfig }

// TranscriptTableReportConfig
//
// See org.apidb.apicommon.model.report.TranscriptTableReporter
type TranscriptTableReportConfig struct{ TableTabularReportConfig }
