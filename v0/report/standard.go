package report

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	errBadStreamStrat = "invalid stream strategy %s"
)

type StreamStrategy string

const (
	StreamStratPaged StreamStrategy = "PAGED_ANSWER"
	StreamStratFile  StreamStrategy = "FILE_BASED"
	StreamStratSmart StreamStrategy = "SMART"
)

func (s StreamStrategy) IsValid() bool {
	switch StreamStrategy(strings.ToUpper(string(s))) {
	case StreamStratPaged:
		return true
	case StreamStratFile:
		return true
	case StreamStratSmart:
		return true
	}
	return false
}

func (s *StreamStrategy) UnmarshalJSON(bytes []byte) error {
	var val string
	if err := json.Unmarshal(bytes, &val); err != nil {
		return err
	}
	tmp := StreamStrategy(val)
	if !tmp.IsValid() {
		return fmt.Errorf(errBadAttType, val)
	}
	*s = tmp
	return nil
}

func (s StreamStrategy) MarshalJSON() ([]byte, error) {
	if !s.IsValid() {
		return nil, fmt.Errorf(errBadAttType, s)
	}
	return json.Marshal(string(s))
}

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

func (s *StandardReportConfig) ValidateReportConfig() error {
	return nil
}

// JsonReportConfig
//
// See org.gusdb.wdk.model.report.reporter.JSONReporter
type JsonReportConfig struct{ StandardReportConfig }

// XmlReportConfig
//
// See org.gusdb.wdk.model.report.reporter.XMLReporter
type XmlReportConfig struct{ StandardReportConfig }
