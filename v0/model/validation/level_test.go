package validation_test

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/validation"
	C "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLevel_IsValid(t *testing.T) {
	C.Convey("Level.IsValid", t, func() {
		C.So(validation.LevelNone.IsValid(), C.ShouldBeTrue)
		C.So(validation.LevelUnspecified.IsValid(), C.ShouldBeTrue)
		C.So(validation.LevelSyntactic.IsValid(), C.ShouldBeTrue)
		C.So(validation.LevelSemantic.IsValid(), C.ShouldBeTrue)
		C.So(validation.LevelRunnable.IsValid(), C.ShouldBeTrue)
		C.So(validation.Level("saturate").IsValid(), C.ShouldBeFalse)
	})
}

func TestLevel_UnmarshalJSON(t *testing.T) {
	C.Convey("Level.UnmarshalJSON", t, func() {
		C.Convey("Invalid json", func() {
			var test validation.Level
			C.So(test.UnmarshalJSON([]byte("^")), C.ShouldNotBeNil)
		})

		C.Convey("Bad Value", func() {
			var test validation.Level
			C.So(test.UnmarshalJSON([]byte(`"blackout"`)), C.ShouldNotBeNil)
		})

		C.Convey("Valid Value", func() {
			var test validation.Level
			C.So(test.UnmarshalJSON([]byte(`"`+validation.LevelSemantic+`"`)), C.ShouldBeNil)
			C.So(test, C.ShouldEqual, validation.LevelSemantic)
		})
	})
}

func TestLevel_MarshalJSON(t *testing.T) {
	C.Convey("Level.MarshalJSON", t, func() {
		C.Convey("Bad Value", func() {
			test := validation.Level("entertainment")
			_, b := test.MarshalJSON()
			C.So(b, C.ShouldNotBeNil)
		})

		C.Convey("Valid Value", func() {
			test := validation.LevelSyntactic
			a, b := test.MarshalJSON()
			C.So(b, C.ShouldBeNil)
			C.So(string(a), C.ShouldEqual, `"SYNTACTIC"`)
		})
	})
}
