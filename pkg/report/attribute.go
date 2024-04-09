package report

// AttributeReportConfig
//
// See org.gusdb.wdk.model.report.AbstractAttributeReporter
type AttributeReportConfig struct{}

func (a *AttributeReportConfig) ValidateReportConfig() error

// EbrcWordCloudAttributeReportConfig
//
// See org.eupathdb.common.model.report.EbrcWordCloudAttributeReporter
type EbrcWordCloudAttributeReportConfig struct{ WordCloudAttributeReportConfig }

// HistogramAttributeReportConfig
//
// See org.gusdb.wdk.model.report.reporter.HistogramAttributeReporter
type HistogramAttributeReportConfig struct{ AttributeReportConfig }

// WordCloudAttributeReportConfig
//
// See org.gusdb.wdk.model.report.reporter.WordCloudAttributeReporter
type WordCloudAttributeReportConfig struct{ AttributeReportConfig }
