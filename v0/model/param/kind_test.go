package param_test

import (
	"encoding/json"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/param"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestKind_IsValid(t *testing.T) {
	convey.Convey("Kind.IsValid", t, func() {
		convey.So(param.KindDate.IsValid(), convey.ShouldBeTrue)
		convey.So(param.KindDateRange.IsValid(), convey.ShouldBeTrue)
		convey.So(param.KindFilter.IsValid(), convey.ShouldBeTrue)
		convey.So(param.KindDataset.IsValid(), convey.ShouldBeTrue)
		convey.So(param.KindAnswer.IsValid(), convey.ShouldBeTrue)
		convey.So(param.KindNumber.IsValid(), convey.ShouldBeTrue)
		convey.So(param.KindNumberRange.IsValid(), convey.ShouldBeTrue)
		convey.So(param.KindString.IsValid(), convey.ShouldBeTrue)
		convey.So(param.KindTimestamp.IsValid(), convey.ShouldBeTrue)
		convey.So(param.KindSingleVocab.IsValid(), convey.ShouldBeTrue)
		convey.So(param.KindMultiVocab.IsValid(), convey.ShouldBeTrue)
		convey.So(param.Kind("foo").IsValid(), convey.ShouldBeFalse)
	})
}

func TestKind_MarshalJSON(t *testing.T) {
	convey.Convey("Kind.MarshalJSON", t, func() {
		var err error
		_, err = json.Marshal(param.KindDate)
		convey.So(err, convey.ShouldBeNil)
		_, err = json.Marshal(param.KindDateRange)
		convey.So(err, convey.ShouldBeNil)
		_, err = json.Marshal(param.KindFilter)
		convey.So(err, convey.ShouldBeNil)
		_, err = json.Marshal(param.KindDataset)
		convey.So(err, convey.ShouldBeNil)
		_, err = json.Marshal(param.KindAnswer)
		convey.So(err, convey.ShouldBeNil)
		_, err = json.Marshal(param.KindNumber)
		convey.So(err, convey.ShouldBeNil)
		_, err = json.Marshal(param.KindNumberRange)
		convey.So(err, convey.ShouldBeNil)
		_, err = json.Marshal(param.KindString)
		convey.So(err, convey.ShouldBeNil)
		_, err = json.Marshal(param.KindTimestamp)
		convey.So(err, convey.ShouldBeNil)
		_, err = json.Marshal(param.KindSingleVocab)
		convey.So(err, convey.ShouldBeNil)
		_, err = json.Marshal(param.KindMultiVocab)
		convey.So(err, convey.ShouldBeNil)
		_, err = json.Marshal(param.Kind("foo"))
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func TestKind_UnmarshalJSON(t *testing.T) {
	convey.Convey("Kind.UnmarshalJSON", t, func() {
		var val param.Kind
		convey.So(json.Unmarshal([]byte(`"`+param.KindDate+`"`), &val), convey.ShouldBeNil)
		convey.So(val, convey.ShouldResemble, param.KindDate)
		convey.So(json.Unmarshal([]byte(`"`+param.KindDateRange+`"`), &val), convey.ShouldBeNil)
		convey.So(val, convey.ShouldResemble, param.KindDateRange)
		convey.So(json.Unmarshal([]byte(`"`+param.KindFilter+`"`), &val), convey.ShouldBeNil)
		convey.So(val, convey.ShouldResemble, param.KindFilter)
		convey.So(json.Unmarshal([]byte(`"`+param.KindDataset+`"`), &val), convey.ShouldBeNil)
		convey.So(val, convey.ShouldResemble, param.KindDataset)
		convey.So(json.Unmarshal([]byte(`"`+param.KindAnswer+`"`), &val), convey.ShouldBeNil)
		convey.So(val, convey.ShouldResemble, param.KindAnswer)
		convey.So(json.Unmarshal([]byte(`"`+param.KindNumber+`"`), &val), convey.ShouldBeNil)
		convey.So(val, convey.ShouldResemble, param.KindNumber)
		convey.So(json.Unmarshal([]byte(`"`+param.KindNumberRange+`"`), &val), convey.ShouldBeNil)
		convey.So(val, convey.ShouldResemble, param.KindNumberRange)
		convey.So(json.Unmarshal([]byte(`"`+param.KindString+`"`), &val), convey.ShouldBeNil)
		convey.So(val, convey.ShouldResemble, param.KindString)
		convey.So(json.Unmarshal([]byte(`"`+param.KindTimestamp+`"`), &val), convey.ShouldBeNil)
		convey.So(val, convey.ShouldResemble, param.KindTimestamp)
		convey.So(json.Unmarshal([]byte(`"`+param.KindSingleVocab+`"`), &val), convey.ShouldBeNil)
		convey.So(val, convey.ShouldResemble, param.KindSingleVocab)
		convey.So(json.Unmarshal([]byte(`"`+param.KindMultiVocab+`"`), &val), convey.ShouldBeNil)
		convey.So(val, convey.ShouldResemble, param.KindMultiVocab)
		convey.So(json.Unmarshal([]byte(`"foo"`), &val), convey.ShouldNotBeNil)
		convey.So(json.Unmarshal([]byte(`{}`), &val), convey.ShouldNotBeNil)
	})
}
