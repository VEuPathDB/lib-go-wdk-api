package report

type Config interface {
	ValidateReportConfig() error
}
