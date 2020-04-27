package common_test

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/common"
	C "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSortDirection_UnmarshalJSON(t *testing.T) {
	C.Convey("SortDirection.UnmarshalJSON", t, func() {
		C.Convey("Invalid json", func() {
			var test common.SortDirection
			C.So(test.UnmarshalJSON([]byte("^")), C.ShouldNotBeNil)
		})

		C.Convey("Bad Value", func() {
			var test common.SortDirection
			C.So(test.UnmarshalJSON([]byte(`"blackout"`)), C.ShouldNotBeNil)
		})

		C.Convey("Valid Value", func() {
			var test common.SortDirection
			C.So(test.UnmarshalJSON([]byte(`"`+common.SortDirectionDesc+`"`)), C.ShouldBeNil)
			C.So(test, C.ShouldEqual, common.SortDirectionDesc)
		})
	})
}

func TestSortDirection_MarshalJSON(t *testing.T) {
	C.Convey("SortDirection.MarshalJSON", t, func() {
		C.Convey("Bad Value", func() {
			test := common.SortDirection("entertainment")
			_, b := test.MarshalJSON()
			C.So(b, C.ShouldNotBeNil)
		})

		C.Convey("Valid Value", func() {
			test := common.SortDirectionAsc
			a, b := test.MarshalJSON()
			C.So(b, C.ShouldBeNil)
			C.So(string(a), C.ShouldEqual, `"ASC"`)
		})
	})
}
