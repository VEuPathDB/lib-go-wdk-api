package answer

import "encoding/json"

// ColumnFilterMap is a mapping from column name to a
// mapping of column tool name to a column tool config.
//
//     colName -> toolName -> config
//
// TODO: make a multitype for the tool config types.
type ColumnFilterMap map[string]map[string]json.RawMessage
