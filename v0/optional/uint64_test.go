package optional_test

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
	C "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewUint64(t *testing.T) {
	C.Convey("NewUint64", t, func() {
		C.So(optional.NewUint64(3).Get(), C.ShouldEqual, 3)
		C.So(optional.NewUint64(7).Exists(), C.ShouldBeTrue)
	})
}

func TestUint64_Clear(t *testing.T) {
	C.Convey("Uint64.Clear", t, func() {
		val := optional.NewUint64(76)
		val.Clear()
		C.So(val.Exists(), C.ShouldBeFalse)
		C.So(func() { val.Get() }, C.ShouldPanic)
	})
}

func TestUint64_Set(t *testing.T) {
	C.Convey("Uint64.Set", t, func() {
		val := optional.Uint64{}
		C.So(val.Exists(), C.ShouldBeFalse)
		val.Set(86)
		C.So(val.Exists(), C.ShouldBeTrue)
		C.So(val.Get(), C.ShouldEqual, 86)
	})
}

func TestUint64_UnmarshalJSON(t *testing.T) {
	C.Convey("Uint64.UnmarshalJSON", t, func() {
		val := optional.Uint64{}
		C.So(val.UnmarshalJSON([]byte(`356`)), C.ShouldBeNil)
		C.So(val.Get(), C.ShouldEqual, 356)
	})
}

func TestUint64_MarshalJSON(t *testing.T) {
	C.Convey("Uint64.MarshalJSON", t, func() {
		val := optional.NewUint64(666)
		a, b := val.MarshalJSON()
		C.So(b, C.ShouldBeNil)
		C.So(string(a), C.ShouldEqual, `666`)
	})

}
