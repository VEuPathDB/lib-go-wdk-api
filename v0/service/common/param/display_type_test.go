package param_test

import (
	"encoding/json"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/common/param"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestDisplayType_IsValid(t *testing.T) {
	convey.Convey("DisplayType.IsValid", t, func() {
		all := param.DisplayTypeValues()
		for _, dt := range all {
			convey.Convey(string(dt), func() {
				convey.So(dt.IsValid(), convey.ShouldBeTrue)
			})
		}
		convey.So(param.DisplayType("foo").IsValid(), convey.ShouldBeFalse)
	})
}

func TestDisplayType_MarshalJSON(t *testing.T) {
	convey.Convey("DisplayType.MarshalJSON", t, func() {
		var err error

		all := param.DisplayTypeValues()

		for _, dt := range all {
			convey.Convey(string(dt), func() {
				_, err = json.Marshal(dt)
				convey.So(err, convey.ShouldBeNil)
			})
		}

		_, err = json.Marshal(param.DisplayType("foo"))
		convey.So(err, convey.ShouldNotBeNil)
	})
}

func TestDisplayType_UnmarshalJSON(t *testing.T) {
	convey.Convey("DisplayType.UnmarshalJSON", t, func() {
		var val param.DisplayType

		all := param.DisplayTypeValues()

		for _, dt := range all {
			convey.Convey(string(dt), func() {
				convey.So(json.Unmarshal([]byte(`"` + dt + `"`), &val), convey.ShouldBeNil)
				convey.So(val, convey.ShouldResemble, dt)
			})
		}
		convey.So(json.Unmarshal([]byte(`"foo"`), &val), convey.ShouldNotBeNil)
		convey.So(json.Unmarshal([]byte(`{}`), &val), convey.ShouldNotBeNil)
	})
}
