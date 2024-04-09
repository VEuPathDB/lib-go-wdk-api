package param

import (
	"encoding/json"
	"errors"
)

type Enum struct {
	Base

	DisplayType      EnumDisplayType `json:"displayType"`
	MaxSelectedCount uint            `json:"maxSelectedCount"`
	MinSelectedCount uint            `json:"minSelectedCount"`
	Vocabulary       Vocabulary      `json:"vocabulary"`

	/* Tree Box Param Fields */
	CountOnlyLeaves *bool `json:"countOnlyLeaves"`
	DepthExpanded   *uint `json:"depthExpanded"`
}

const (
	EnumDisplayTypeSelect    EnumDisplayType = "select"
	EnumDisplayTypeCheckBox  EnumDisplayType = "checkBox"
	EnumDisplayTypeTreeBox   EnumDisplayType = "treeBox"
	EnumDisplayTypeTypeAhead EnumDisplayType = "typeAhead"
)

type EnumDisplayType string

func EnumDisplayTypeValues() []EnumDisplayType {
	return []EnumDisplayType{EnumDisplayTypeSelect, EnumDisplayTypeCheckBox, EnumDisplayTypeTreeBox, EnumDisplayTypeTypeAhead}
}

const (
	vocabKindNone  uint8 = 0
	vocabKindTable uint8 = 1
	vocabKindTree  uint8 = 2
)

type TreeVocab struct {
	Data     TreeVocabData `json:"data"`
	Children []TreeVocab   `json:"children,omitempty"`
}

type TreeVocabData struct {
	Term    string `json:"term"`
	Display string `json:"display"`
}

type TableVocab [][3]string

// Vocabulary could be a TableVocab, a TreeVocab, or empty.
type Vocabulary struct {
	value any
	kind  uint8
}

// Clear removes the value contained in this Vocabulary and returns it to an
// empty state.
func (v *Vocabulary) Clear() {
	v.value = nil
	v.kind = vocabKindNone
}

// Exists returns whether this Vocabulary struct contains any value.
func (v *Vocabulary) Exists() bool {
	return v.kind != vocabKindNone
}

// GetAsTable returns the internal value of this Vocabulary as a TableVocab.
//
// If the internal value exists and is not a TableVocab, this method will panic.
// Verify that the contained value is a TableVocab using the IsTableVocab
// method.
//
// If this Vocabulary is empty, returns nil.
func (v *Vocabulary) GetAsTable() TableVocab {
	if v.kind != vocabKindTable {
		panic("incorrectly attempted to unwrap a param.Vocabulary instance as a table vocabulary")
	}

	return v.value.(TableVocab)
}

// GetAsTree returns the internal value of this Vocabulary as a TreeVocab.
//
// If the internal value exists and is not a TreeVocab, this method will panic.
// Verify that the contained value is a TreeVocab using the IsTableVocab method.
//
// If this Vocabulary is empty, returns nil.
//
// The pointer returned by this method cannot be used to modify the internal
// state of this Vocabulary, if it is desired to update the internal tree, call
// SetTreeVocab with the updated tree.
func (v *Vocabulary) GetAsTree() *TreeVocab {
	if v.kind != vocabKindTree {
		panic("incorrectly attempted to unwrap a param.Vocabulary instance as a tree vocabulary")
	}

	return v.value.(*TreeVocab)
}

// IsTableVocab returns whether this Vocabulary contains a TableVocab value.
func (v *Vocabulary) IsTableVocab() bool {
	return v.kind == vocabKindTable
}

// IsTreeVocab returns whether this Vocabulary contains a TreeVocab value.
func (v *Vocabulary) IsTreeVocab() bool {
	return v.kind == vocabKindTree
}

func (v Vocabulary) MarshalJSON() ([]byte, error) {
	switch v.kind {
	case vocabKindNone:
		return []byte("null"), nil
	case vocabKindTable:
		return json.Marshal(v.GetAsTable())
	case vocabKindTree:
		return json.Marshal(v.GetAsTree())
	default:
		panic("illegal state")
	}
}

// SetTableVocab sets this Vocabulary's value to the given TableVocab and
// configures the internal state to reflect the value type.
func (v *Vocabulary) SetTableVocab(val TableVocab) {
	v.value = val
	v.kind = vocabKindTable
}

// SetTreeVocab sets this Vocabulary's value to the given TreeVocab and
// configures the internal state to reflect the value type.
func (v *Vocabulary) SetTreeVocab(val TreeVocab) {
	v.value = &val
	v.kind = vocabKindTree
}

func (v *Vocabulary) UnmarshalJSON(bytes []byte) error {
	switch bytes[0] {
	case '[':
		var tmp TableVocab
		if err := json.Unmarshal(bytes, &tmp); err != nil {
			return err
		}
		v.SetTableVocab(tmp)

	case '{':
		var tmp TreeVocab
		if err := json.Unmarshal(bytes, &tmp); err != nil {
			return err
		}
		v.SetTreeVocab(tmp)

	default:
		return errors.New("unexpected JSON vocabulary value")
	}

	return nil
}
