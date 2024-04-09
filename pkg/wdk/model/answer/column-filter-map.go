package answer

import "encoding/json"

// ColumnFilterMap defines a mapping from column name to a mapping of column
// tool name to a column tool config.
//
//   colName -> toolName -> config
type ColumnFilterMap map[string]map[string]json.RawMessage
