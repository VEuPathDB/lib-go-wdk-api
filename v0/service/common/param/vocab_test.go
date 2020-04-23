package param_test

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/common/param"
	"testing"

	C "github.com/smartystreets/goconvey/convey"
)

func TestVocabulary_Exists(t *testing.T) {
	C.Convey("Vocabulary.Exists", t, func() {
		C.Convey("initial state", func() {
			var test param.Vocabulary
			C.So(test.Exists(), C.ShouldBeFalse)
		})
		C.Convey("after setting table vocab", func() {
			var test param.Vocabulary
			C.So(test.Exists(), C.ShouldBeFalse)
			test.SetTableVocab(param.TableVocab{})
			C.So(test.Exists(), C.ShouldBeTrue)
		})
		C.Convey("after setting tree vocab", func() {
			var test param.Vocabulary
			C.So(test.Exists(), C.ShouldBeFalse)
			test.SetTreeVocab(param.TreeVocab{})
			C.So(test.Exists(), C.ShouldBeTrue)
		})
		C.Convey("after clear", func() {
			var test param.Vocabulary
			C.So(test.Exists(), C.ShouldBeFalse)
			test.SetTreeVocab(param.TreeVocab{})
			C.So(test.Exists(), C.ShouldBeTrue)
			test.Clear()
			C.So(test.Exists(), C.ShouldBeFalse)
		})
	})
}

func TestVocabulary_IsTableVocab(t *testing.T) {
	C.Convey("Vocabulary.IsTableVocab", t, func() {
		C.Convey("initial state", func() {
			var test param.Vocabulary
			C.So(test.IsTableVocab(), C.ShouldBeFalse)
		})
		C.Convey("after setting table vocab", func() {
			var test param.Vocabulary
			test.SetTableVocab(param.TableVocab{})
			C.So(test.IsTableVocab(), C.ShouldBeTrue)
		})
		C.Convey("after setting tree vocab", func() {
			var test param.Vocabulary
			test.SetTreeVocab(param.TreeVocab{})
			C.So(test.IsTableVocab(), C.ShouldBeFalse)
		})
		C.Convey("after clear", func() {
			var test param.Vocabulary
			test.SetTableVocab(param.TableVocab{})
			test.Clear()
			C.So(test.IsTableVocab(), C.ShouldBeFalse)
		})
		C.Convey("after setting tree then table vocab", func() {
			var test param.Vocabulary
			test.SetTreeVocab(param.TreeVocab{})
			test.SetTableVocab(param.TableVocab{})
			C.So(test.IsTableVocab(), C.ShouldBeTrue)
		})
		C.Convey("after setting table then tree vocab", func() {
			var test param.Vocabulary
			test.SetTableVocab(param.TableVocab{})
			test.SetTreeVocab(param.TreeVocab{})
			C.So(test.IsTableVocab(), C.ShouldBeFalse)
		})
	})
}

func TestVocabulary_IsTreeVocab(t *testing.T) {
	C.Convey("Vocabulary.IsTreeVocab", t, func() {
		C.Convey("initial state", func() {
			var test param.Vocabulary
			C.So(test.IsTreeVocab(), C.ShouldBeFalse)
		})
		C.Convey("after setting table vocab", func() {
			var test param.Vocabulary
			test.SetTableVocab(param.TableVocab{})
			C.So(test.IsTreeVocab(), C.ShouldBeFalse)
		})
		C.Convey("after setting tree vocab", func() {
			var test param.Vocabulary
			test.SetTreeVocab(param.TreeVocab{})
			C.So(test.IsTreeVocab(), C.ShouldBeTrue)
		})
		C.Convey("after clear", func() {
			var test param.Vocabulary
			test.SetTreeVocab(param.TreeVocab{})
			test.Clear()
			C.So(test.IsTreeVocab(), C.ShouldBeFalse)
		})
		C.Convey("after setting tree then table vocab", func() {
			var test param.Vocabulary
			test.SetTreeVocab(param.TreeVocab{})
			test.SetTableVocab(param.TableVocab{})
			C.So(test.IsTreeVocab(), C.ShouldBeFalse)
		})
		C.Convey("after setting table then tree vocab", func() {
			var test param.Vocabulary
			test.SetTableVocab(param.TableVocab{})
			test.SetTreeVocab(param.TreeVocab{})
			C.So(test.IsTreeVocab(), C.ShouldBeTrue)
		})
	})
}

func TestVocabulary_GetAsTable(t *testing.T) {
	C.Convey("Vocabulary.GetAsTable", t, func() {
		C.Convey("initial state", func() {
			var test param.Vocabulary
			C.So(test.GetAsTable(), C.ShouldBeNil)
		})
		C.Convey("after setting table vocab", func() {
			var test param.Vocabulary
			input := param.TableVocab{ [3]string{"a, b", "c", "d, e"} }
			test.SetTableVocab(input)
			C.So(test.GetAsTable(), C.ShouldResemble, input)
		})
		C.Convey("after setting tree vocab", func() {
			var test param.Vocabulary
			test.SetTreeVocab(param.TreeVocab{})
			C.So(func() {test.GetAsTable()}, C.ShouldPanic)
		})
		C.Convey("after table clear", func() {
			var test param.Vocabulary
			test.SetTableVocab(param.TableVocab{ [3]string{"a, b", "c", "d, e"} })
			test.Clear()
			C.So(test.GetAsTable(), C.ShouldBeNil)
		})
		C.Convey("after tree clear", func() {
			var test param.Vocabulary
			test.SetTreeVocab(param.TreeVocab{})
			test.Clear()
			C.So(test.GetAsTable(), C.ShouldBeNil)
		})
		C.Convey("after setting tree then table vocab", func() {
			var test param.Vocabulary
			input := param.TableVocab{ [3]string{"a, b", "c", "d, e"} }
			test.SetTreeVocab(param.TreeVocab{})
			test.SetTableVocab(input)
			C.So(test.GetAsTable(), C.ShouldResemble, input)
		})
		C.Convey("after setting table then tree vocab", func() {
			var test param.Vocabulary
			input := param.TableVocab{ [3]string{"a, b", "c", "d, e"} }
			test.SetTableVocab(input)
			test.SetTreeVocab(param.TreeVocab{})
			C.So(func() {test.GetAsTable()}, C.ShouldPanic)
		})
	})
}

func TestVocabulary_GetAsTree(t *testing.T) {
	C.Convey("Vocabulary.GetAsTree", t, func() {
		C.Convey("initial state", func() {
			var test param.Vocabulary
			C.So(test.GetAsTree(), C.ShouldBeNil)
		})
		C.Convey("after setting table vocab", func() {
			var test param.Vocabulary
			input := param.TableVocab{ [3]string{"a, b", "c", "d, e"} }
			test.SetTableVocab(input)
			C.So(func() {test.GetAsTree()}, C.ShouldPanic)
		})
		C.Convey("after setting tree vocab", func() {
			var test param.Vocabulary
			input := param.TreeVocab{
				Data: param.TreeVocabData{Term: "a", Display: "B"},
				Children: []param.TreeVocab{{
					Data: param.TreeVocabData{Term: "c", Display: "D"}}}}
			test.SetTreeVocab(input)
			C.So(test.GetAsTree(), C.ShouldResemble, &input)
		})
		C.Convey("after table clear", func() {
			var test param.Vocabulary
			test.SetTableVocab(param.TableVocab{ [3]string{"a, b", "c", "d, e"} })
			test.Clear()
			C.So(test.GetAsTree(), C.ShouldBeNil)
		})
		C.Convey("after tree clear", func() {
			var test param.Vocabulary
			input := param.TreeVocab{
				Data: param.TreeVocabData{Term: "a", Display: "B"},
				Children: []param.TreeVocab{{
					Data: param.TreeVocabData{Term: "c", Display: "D"}}}}
			test.SetTreeVocab(input)
			test.Clear()
			C.So(test.GetAsTree(), C.ShouldBeNil)
		})
		C.Convey("after setting tree then table vocab", func() {
			var test param.Vocabulary
			input2 := param.TreeVocab{
				Data: param.TreeVocabData{Term: "a", Display: "B"},
				Children: []param.TreeVocab{{
					Data: param.TreeVocabData{Term: "c", Display: "D"}}}}
			test.SetTreeVocab(input2)
			input1 := param.TableVocab{ [3]string{"a, b", "c", "d, e"} }
			test.SetTableVocab(input1)
			C.So(func() {test.GetAsTree()}, C.ShouldPanic)
		})
		C.Convey("after setting table then tree vocab", func() {
			var test param.Vocabulary
			input1 := param.TableVocab{[3]string{"a, b", "c", "d, e"} }
			test.SetTableVocab(input1)
			input2 := param.TreeVocab{
				Data: param.TreeVocabData{Term: "a", Display: "B"},
				Children: []param.TreeVocab{{
					Data: param.TreeVocabData{Term: "c", Display: "D"}}}}
			test.SetTreeVocab(input2)
			C.So(test.GetAsTree(), C.ShouldResemble, &input2)
		})
	})
}

func TestVocabulary_MarshalJSON(t *testing.T) {
	C.Convey("Vocabulary.MarshalJSON", t, func() {
		C.Convey("initial state", func() {
			var test param.Vocabulary
			a, b := test.MarshalJSON()
			C.So(b, C.ShouldBeNil)
			C.So(a, C.ShouldResemble, []byte("null"))
		})
		C.Convey("after setting table vocab", func() {
			var test param.Vocabulary
			test.SetTableVocab(param.TableVocab{ [3]string{"a, b", "c", "d, e"} })

			a, b := test.MarshalJSON()
			C.So(b, C.ShouldBeNil)
			C.So(a, C.ShouldResemble, []byte(`[["a, b","c","d, e"]]`))
		})
		C.Convey("after setting tree vocab", func() {
			var test param.Vocabulary
			input := param.TreeVocab{
				Data: param.TreeVocabData{Term: "a", Display: "B"},
				Children: []param.TreeVocab{{
					Data: param.TreeVocabData{Term: "c", Display: "D"}}}}
			test.SetTreeVocab(input)

			a, b := test.MarshalJSON()
			C.So(b, C.ShouldBeNil)
			C.So(string(a), C.ShouldResemble, `{"data":{"term":"a","display":"B"},"children":[{"data":{"term":"c","display":"D"}}]}`)
		})
		C.Convey("after table clear", func() {
			var test param.Vocabulary
			test.SetTableVocab(param.TableVocab{ [3]string{"a, b", "c", "d, e"} })
			test.Clear()

			a, b := test.MarshalJSON()
			C.So(b, C.ShouldBeNil)
			C.So(a, C.ShouldResemble, []byte("null"))
		})
		C.Convey("after tree clear", func() {
			var test param.Vocabulary
			input := param.TreeVocab{
				Data: param.TreeVocabData{Term: "a", Display: "B"},
				Children: []param.TreeVocab{{
					Data: param.TreeVocabData{Term: "c", Display: "D"}}}}
			test.SetTreeVocab(input)
			test.Clear()

			a, b := test.MarshalJSON()
			C.So(b, C.ShouldBeNil)
			C.So(a, C.ShouldResemble, []byte("null"))
		})
	})
}

func TestVocabulary_UnmarshalJSON(t *testing.T) {
	C.Convey("Vocabulary.UnmarshalJSON", t, func() {
		C.Convey("unmarshal null", func() {
			var test param.Vocabulary
			C.So(test.UnmarshalJSON([]byte("null")), C.ShouldBeNil)
			C.So(test.Exists(), C.ShouldBeFalse)
		})

		C.Convey("unmarshal invalid json", func() {
			tests := [][]byte{[]byte("nigel"), []byte("{arf}"), []byte("[boo]")}

			for _, test := range tests {
				var voc param.Vocabulary
				C.So(voc.UnmarshalJSON(test), C.ShouldNotBeNil)
				C.So(voc.Exists(), C.ShouldBeFalse)
			}
		})

		C.Convey("unmarshal incorrect type", func() {
			var test param.Vocabulary
			C.So(test.UnmarshalJSON([]byte(`"hello"`)), C.ShouldNotBeNil)
			C.So(test.Exists(), C.ShouldBeFalse)
		})

		C.Convey("unmarshal table", func() {
			var test param.Vocabulary
			value := `[["a","b","c"],["d","e","f"]]`

			C.So(test.UnmarshalJSON([]byte(value)), C.ShouldBeNil)
			C.So(test.Exists(), C.ShouldBeTrue)
			C.So(test.IsTableVocab(), C.ShouldBeTrue)
			C.So(test.IsTreeVocab(), C.ShouldBeFalse)
			C.So(test.GetAsTable(), C.ShouldResemble, param.TableVocab{
				[3]string{"a", "b", "c"},
				[3]string{"d", "e", "f"},
			})
		})

		C.Convey("unmarshal tree", func() {
			var test param.Vocabulary
			value := `{"data":{"term":"a","display":"apple"},"children":[{"data":{"term":"b","display":"blueberry"}},{"data":{"term":"c","display":"cranberry"}}]}`

			C.So(test.UnmarshalJSON([]byte(value)), C.ShouldBeNil)
			C.So(test.Exists(), C.ShouldBeTrue)
			C.So(test.IsTableVocab(), C.ShouldBeFalse)
			C.So(test.IsTreeVocab(), C.ShouldBeTrue)
			C.So(test.GetAsTree(), C.ShouldResemble, &param.TreeVocab{
				Data: param.TreeVocabData{
					Term:    "a",
					Display: "apple",
				},
				Children: []param.TreeVocab{
					{Data: param.TreeVocabData{Term: "b", Display: "blueberry"}},
					{Data: param.TreeVocabData{Term: "c", Display: "cranberry"}},
				},
			})
		})
	})
}