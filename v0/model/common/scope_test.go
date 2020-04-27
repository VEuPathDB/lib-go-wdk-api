package common_test

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/common"
	C "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestScope_IsValid(t *testing.T) {
	C.Convey("Scope.IsValid", t, func() {
		C.So(common.ScopeRecord.IsValid(), C.ShouldBeTrue)
		C.So(common.ScopeResult.IsValid(), C.ShouldBeTrue)
		C.So(common.Scope("saturate").IsValid(), C.ShouldBeFalse)
	})
}

func TestScope_UnmarshalJSON(t *testing.T) {
	C.Convey("Scope.UnmarshalJSON", t, func() {
		C.Convey("Invalid json", func() {
			var test common.Scope
			C.So(test.UnmarshalJSON([]byte("^")), C.ShouldNotBeNil)
		})

		C.Convey("Bad Value", func() {
			var test common.Scope
			C.So(test.UnmarshalJSON([]byte(`"blackout"`)), C.ShouldNotBeNil)
		})

		C.Convey("Valid Value", func() {
			var test common.Scope
			C.So(test.UnmarshalJSON([]byte(`"`+common.ScopeResult+`"`)), C.ShouldBeNil)
			C.So(test, C.ShouldEqual, common.ScopeResult)
		})
	})
}

func TestScope_MarshalJSON(t *testing.T) {
	C.Convey("Scope.MarshalJSON", t, func() {
		C.Convey("Bad Value", func() {
			test := common.Scope("entertainment")
			_, b := test.MarshalJSON()
			C.So(b, C.ShouldNotBeNil)
		})

		C.Convey("Valid Value", func() {
			test := common.ScopeRecord
			a, b := test.MarshalJSON()
			C.So(b, C.ShouldBeNil)
			C.So(string(a), C.ShouldEqual, `"record"`)
		})
	})
}
