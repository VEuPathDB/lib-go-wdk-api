package common_test

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/common"
	C "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAllMarker_IsValid(t *testing.T) {
	C.Convey("AllMarker.IsValid", t, func() {
		C.So(common.AllAttributes.IsValid(), C.ShouldBeTrue)
		C.So(common.AllTables.IsValid(), C.ShouldBeTrue)
		C.So(common.AllMarker("saturate").IsValid(), C.ShouldBeFalse)
	})
}

func TestAllMarker_UnmarshalJSON(t *testing.T) {
	C.Convey("AllMarker.UnmarshalJSON", t, func() {
		C.Convey("Invalid json", func() {
			var test common.AllMarker
			C.So(test.UnmarshalJSON([]byte("^")), C.ShouldNotBeNil)
		})

		C.Convey("Bad Value", func() {
			var test common.AllMarker
			C.So(test.UnmarshalJSON([]byte(`"blackout"`)), C.ShouldNotBeNil)
		})

		C.Convey("Valid Value", func() {
			var test common.AllMarker
			C.So(test.UnmarshalJSON([]byte(`"`+common.AllTables+`"`)), C.ShouldBeNil)
			C.So(test, C.ShouldEqual, common.AllTables)
		})
	})
}

func TestAllMarker_MarshalJSON(t *testing.T) {
	C.Convey("AllMarker.MarshalJSON", t, func() {
		C.Convey("Bad Value", func() {
			test := common.AllMarker("entertainment")
			_, b := test.MarshalJSON()
			C.So(b, C.ShouldNotBeNil)
		})

		C.Convey("Valid Value", func() {
			test := common.AllAttributes
			a, b := test.MarshalJSON()
			C.So(b, C.ShouldBeNil)
			C.So(string(a), C.ShouldEqual, `"__ALL_ATTRIBUTES__"`)
		})
	})
}
