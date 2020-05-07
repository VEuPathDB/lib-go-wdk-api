package attribute_test

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/attribute"
	C "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAllMarker_IsValid(t *testing.T) {
	C.Convey("Preset.IsValid", t, func() {
		C.So(attribute.PresetAllAttributes.IsValid(), C.ShouldBeTrue)
		C.So(attribute.PresetDefaultAttributes.IsValid(), C.ShouldBeTrue)
		C.So(attribute.Preset("saturate").IsValid(), C.ShouldBeFalse)
	})
}

func TestAllMarker_UnmarshalJSON(t *testing.T) {
	C.Convey("Preset.UnmarshalJSON", t, func() {
		C.Convey("Invalid json", func() {
			var test attribute.Preset
			C.So(test.UnmarshalJSON([]byte("^")), C.ShouldNotBeNil)
		})

		C.Convey("Bad Value", func() {
			var test attribute.Preset
			C.So(test.UnmarshalJSON([]byte(`"blackout"`)), C.ShouldNotBeNil)
		})

		C.Convey("Valid Value", func() {
			var test attribute.Preset
			C.So(test.UnmarshalJSON([]byte(`"`+attribute.PresetDefaultAttributes+`"`)), C.ShouldBeNil)
			C.So(test, C.ShouldEqual, attribute.PresetDefaultAttributes)
		})
	})
}

func TestAllMarker_MarshalJSON(t *testing.T) {
	C.Convey("Preset.MarshalJSON", t, func() {
		C.Convey("Bad Value", func() {
			test := attribute.Preset("entertainment")
			_, b := test.MarshalJSON()
			C.So(b, C.ShouldNotBeNil)
		})

		C.Convey("Valid Value", func() {
			test := attribute.PresetAllAttributes
			a, b := test.MarshalJSON()
			C.So(b, C.ShouldBeNil)
			C.So(string(a), C.ShouldEqual, `"__ALL_ATTRIBUTES__"`)
		})
	})
}
