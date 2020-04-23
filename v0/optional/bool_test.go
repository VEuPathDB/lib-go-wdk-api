package optional_test

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
	C "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewBool(t *testing.T) {
	C.Convey("NewBool", t, func() {
		C.So(optional.NewBool(true).Get(), C.ShouldEqual, true)
		C.So(optional.NewBool(false).Exists(), C.ShouldBeTrue)
	})
}

func TestBool_Clear(t *testing.T) {
	C.Convey("Bool.Clear", t, func() {
		val := optional.NewBool(true)
		val.Clear()
		C.So(val.Exists(), C.ShouldBeFalse)
		C.So(func() {val.Get()}, C.ShouldPanic)
	})
}

func TestBool_Set(t *testing.T) {
	C.Convey("Bool.Set", t, func() {
		val := optional.Bool{}
		C.So(val.Exists(), C.ShouldBeFalse)
		val.Set(false)
		C.So(val.Exists(), C.ShouldBeTrue)
		C.So(val.Get(), C.ShouldEqual, false)
	})
}

func TestBool_UnmarshalJSON(t *testing.T) {
	C.Convey("Bool.UnmarshalJSON", t, func() {
		val := optional.Bool{}
		C.So(val.UnmarshalJSON([]byte(`true`)), C.ShouldBeNil)
		C.So(val.Get(), C.ShouldEqual, true)
	})
}

func TestBool_MarshalJSON(t *testing.T) {
	C.Convey("Bool.MarshalJSON", t, func() {
		val := optional.NewBool(false)
		a, b := val.MarshalJSON()
		C.So(b, C.ShouldBeNil)
		C.So(string(a), C.ShouldEqual, `false`)
	})

}
