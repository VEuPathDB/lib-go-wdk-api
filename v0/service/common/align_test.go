package common_test

import (
	"encoding/json"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/common"
	"testing"

	C "github.com/smartystreets/goconvey/convey"
)

func TestAlign_IsValid(t *testing.T) {
	C.Convey("Align.IsValid", t, func() {
		C.So(common.AlignCenter.IsValid(), C.ShouldBeTrue)
		C.So(common.AlignLeft.IsValid(), C.ShouldBeTrue)
		C.So(common.AlignRight.IsValid(), C.ShouldBeTrue)
		C.So(common.Align("foo").IsValid(), C.ShouldBeFalse)
	})
}

func TestAlign_UnmarshalJSON(t *testing.T) {
	C.Convey("Align.UnmarshalJSON", t, func() {
		C.Convey("When source value is a known Align value", func() {
			C.Convey("Should succeed", func() {
				var a, b, c common.Align

				C.So(json.Unmarshal([]byte(`"center"`), &a), C.ShouldBeNil)
				C.So(a, C.ShouldResemble, common.AlignCenter)

				C.So(json.Unmarshal([]byte(`"left"`), &b), C.ShouldBeNil)
				C.So(b, C.ShouldResemble, common.AlignLeft)

				C.So(json.Unmarshal([]byte(`"right"`), &c), C.ShouldBeNil)
				C.So(c, C.ShouldResemble, common.AlignRight)
			})
		})

		C.Convey("When source value is not a known Align value", func() {
			C.Convey("Should Fail", func() {
				var a common.Align

				C.So(json.Unmarshal([]byte(`"something"`), &a), C.ShouldNotBeNil)
				C.So(json.Unmarshal([]byte(`{}`), &a), C.ShouldNotBeNil)
			})
		})
	})
}

func TestOptionalAlign(t *testing.T) {
	C.Convey("OptionalAlign", t, func() {
		var oa common.OptionalAlign
		C.So(oa.Exists(), C.ShouldBeFalse)
		C.So(func() {oa.Get()}, C.ShouldPanic)

		oa.Set(common.AlignLeft)
		C.So(oa.Exists(), C.ShouldBeTrue)
		C.So(oa.Get(), C.ShouldResemble, common.AlignLeft)

		oa.Clear()
		C.So(oa.Exists(), C.ShouldBeFalse)
		C.So(func() {oa.Get()}, C.ShouldPanic)

		C.So(oa.UnmarshalJSON([]byte(`"right"`)), C.ShouldBeNil)
		C.So(oa.Exists(), C.ShouldBeTrue)
		C.So(oa.Get(), C.ShouldResemble, common.AlignRight)
	})
}