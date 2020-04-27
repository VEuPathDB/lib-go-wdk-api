package common_test

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/common"
	C "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewOptionalProperties(t *testing.T) {
	C.Convey("NewOptionalProperties", t, func() {
		val := common.Properties{"hello": {"hi"}}
		C.So(common.NewOptionalProperties(val).Get(), C.ShouldResemble, val)
		C.So(common.NewOptionalProperties(val).Exists(), C.ShouldBeTrue)
	})
}

func TestProperties_Clear(t *testing.T) {
	C.Convey("Properties.Clear", t, func() {
		raw := common.Properties{"fallout": {"76"}}
		val := common.NewOptionalProperties(raw)
		val.Clear()
		C.So(val.Exists(), C.ShouldBeFalse)
		C.So(val.Get(), C.ShouldBeNil)
	})
}

func TestProperties_Set(t *testing.T) {
	C.Convey("Properties.Set", t, func() {
		raw := common.Properties{"humpty": {"dumpty"}}
		val := common.OptionalProperties{}
		C.So(val.Exists(), C.ShouldBeFalse)
		val.Set(raw)
		C.So(val.Exists(), C.ShouldBeTrue)
		C.So(val.Get(), C.ShouldResemble, raw)
	})
}

func TestProperties_UnmarshalJSON(t *testing.T) {
	C.Convey("Properties.UnmarshalJSON", t, func() {
		val := common.OptionalProperties{}
		C.So(val.UnmarshalJSON([]byte(`{"good": ["morning"]}`)), C.ShouldBeNil)
		C.So(val.Get(), C.ShouldResemble, common.Properties{"good": {"morning"}})
	})
}

func TestProperties_MarshalJSON(t *testing.T) {
	C.Convey("Properties.MarshalJSON", t, func() {
		val := common.NewOptionalProperties(common.Properties{"good": {"bye"}})
		a, b := val.MarshalJSON()
		C.So(b, C.ShouldBeNil)
		C.So(string(a), C.ShouldEqual, `{"good":["bye"]}`)
	})
}
