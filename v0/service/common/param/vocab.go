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

func (v *Vocabulary) Exists() bool {
	return v.value != nil
}

func (v *Vocabulary) Clear() {
	v.value = nil
	v.isTree = false
	v.isTable = false
}

func (v *Vocabulary) SetTreeVocab(val TreeVocab) {
	v.value = val
	v.isTree = true
	v.isTable = false
}

func (v *Vocabulary) SetTableVocab(val TableVocab) {
	v.value = val
	v.isTree = false
	v.isTable = true
}

func (v *Vocabulary) IsTreeVocab() bool {
	return v.isTree
}

func (v *Vocabulary) IsTableVocab() bool {
	return v.isTable
}

func (v *Vocabulary) GetAsTable() TableVocab {
	if v.value == nil {
		return nil
	}

	if !v.isTable {
		panic("cannot cast TreeVocab as TableVocab")
	}

	return v.value.(TableVocab)
}

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
		v.isTree = false
	case '{':
		var tmp TreeVocab
		if err := json.Unmarshal(bytes, &tmp); err != nil {
			return err
		}
		v.value = tmp
		v.isTree = true
	case 'n':
		if len(bytes) == 4 && string(bytes) == "null" {
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
