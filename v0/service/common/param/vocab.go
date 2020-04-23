package param

import (
	"encoding/json"
	"errors"
)

// Vocabulary could be a TableVocab, a TreeVocab, or empty.
type Vocabulary struct {
	value   interface{}
	isTree  bool
	isTable bool
}

// Exists returns whether this Vocabulary struct contains
// any value.
func (v *Vocabulary) Exists() bool {
	return v.value != nil
}

// Clear removes the value contained in this Vocabulary and
// returns it to an empty state.
func (v *Vocabulary) Clear() {
	v.value = nil
	v.isTree = false
	v.isTable = false
}

// SetTreeVocab sets this Vocabulary's value to the given
// TreeVocab and configures the internal state to reflect
// the value type.
func (v *Vocabulary) SetTreeVocab(val TreeVocab) {
	v.value = val
	v.isTree = true
	v.isTable = false
}

// SetTreeVocab sets this Vocabulary's value to the given
// TableVocab and configures the internal state to reflect
// the value type.
func (v *Vocabulary) SetTableVocab(val TableVocab) {
	v.value = val
	v.isTree = false
	v.isTable = true
}

// IsTreeVocab returns whether or not this Vocabulary
// contains a TreeVocab value.
func (v *Vocabulary) IsTreeVocab() bool {
	return v.isTree
}

// IsTreeVocab returns whether or not this Vocabulary
// contains a TableVocab value.
func (v *Vocabulary) IsTableVocab() bool {
	return v.isTable
}

// GetAsTable returns the internal value of this Vocabulary
// as a TableVocab.
//
// If the internal value exists and is not a TableVocab,
// this method will panic.  Verify that the contained value
// is a TableVocab using the IsTableVocab method.
//
// If this Vocabulary is empty, returns nil.
func (v *Vocabulary) GetAsTable() TableVocab {
	if v.value == nil {
		return nil
	}

	if !v.isTable {
		panic("cannot cast TreeVocab as TableVocab")
	}

	return v.value.(TableVocab)
}

// GetAsTree returns the internal value of this Vocabulary
// as a TreeVocab.
//
// If the internal value exists and is not a TreeVocab,
// this method will panic.  Verify that the contained value
// is a TreeVocab using the IsTableVocab method.
//
// If this Vocabulary is empty, returns nil.
//
// The pointer returned by this method cannot be used to
// modify the internal state of this Vocabulary, if it is
// desired to update the internal tree, call SetTreeVocab
// with the updated tree.
func (v *Vocabulary) GetAsTree() *TreeVocab {
	if v.value == nil {
		return nil
	}

	if !v.isTree {
		panic("cannot cast TableVocab as TreeVocab")
	}

	tmp := v.value.(TreeVocab)
	return &tmp
}

func (v *Vocabulary) UnmarshalJSON(bytes []byte) error {
	if len(bytes) < 2 || (bytes[0] != '[' && bytes[0] != '{' && bytes[0] != 'n') {
		return errors.New("vocabulary input must be an array or object")
	}
	switch bytes[0] {
	case '[':
		var tmp TableVocab
		if err := json.Unmarshal(bytes, &tmp); err != nil {
			return err
		}
		v.value = tmp
		v.isTable = true
		v.isTree = false
	case '{':
		var tmp TreeVocab
		if err := json.Unmarshal(bytes, &tmp); err != nil {
			return err
		}
		v.value = tmp
		v.isTree = true
		v.isTable = false
	case 'n':
		if len(bytes) == 4 && string(bytes) == "null" {
			v.value = nil
			v.isTree = false
			v.isTable = false
			return nil
		}
		return errors.New("invalid json input starting with 'n'")
	}
	return nil
}

func (v Vocabulary) MarshalJSON() ([]byte, error) {
	if v.isTree {
		return json.Marshal(v.value.(TreeVocab))
	} else if v.isTable {
		return json.Marshal(v.value.(TableVocab))
	} else {
		return []byte{'n', 'u', 'l', 'l'}, nil
	}
}

type TableVocab [][3]string

type TreeVocab struct {
	Data     TreeVocabData `json:"data"`
	Children []TreeVocab   `json:"children,omitempty"`
}

type TreeVocabData struct {
	Term    string `json:"term"`
	Display string `json:"display"`
}
