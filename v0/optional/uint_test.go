package optional_test

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
	C "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewUint(t *testing.T) {
	C.Convey("NewUint", t, func() {
		C.So(optional.NewUint(3).Get(), C.ShouldEqual, 3)
		C.So(optional.NewUint(7).Exists(), C.ShouldBeTrue)
	})
}

func TestUint_Clear(t *testing.T) {
	C.Convey("Uint.Clear", t, func() {
		val := optional.NewUint(76)
		val.Clear()
		C.So(val.Exists(), C.ShouldBeFalse)
		C.So(func() {val.Get()}, C.ShouldPanic)
	})
}

func TestUint_Set(t *testing.T) {
	C.Convey("Uint.Set", t, func() {
		val := optional.Uint{}
		C.So(val.Exists(), C.ShouldBeFalse)
		val.Set(86)
		C.So(val.Exists(), C.ShouldBeTrue)
		C.So(val.Get(), C.ShouldEqual, 86)
	})
}

func TestUint_UnmarshalJSON(t *testing.T) {
	C.Convey("Uint.UnmarshalJSON", t, func() {
		val := optional.Uint{}
		C.So(val.UnmarshalJSON([]byte(`356`)), C.ShouldBeNil)
		C.So(val.Get(), C.ShouldEqual, 356)
	})
}

func TestUint_MarshalJSON(t *testing.T) {
	C.Convey("Uint.MarshalJSON", t, func() {
		val := optional.NewUint(666)
		a, b := val.MarshalJSON()
		C.So(b, C.ShouldBeNil)
		C.So(string(a), C.ShouldEqual, `666`)
	})

}
