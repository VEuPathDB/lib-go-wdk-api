package optional_test

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
	C "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewString(t *testing.T) {
	C.Convey("NewString", t, func() {
		C.So(optional.NewString("foo").Get(), C.ShouldEqual, "foo")
		C.So(optional.NewString("bar").Exists(), C.ShouldBeTrue)
	})
}

func TestString_Clear(t *testing.T) {
	C.Convey("String.Clear", t, func() {
		val := optional.NewString("flump")
		val.Clear()
		C.So(val.Exists(), C.ShouldBeFalse)
		C.So(func() {val.Get()}, C.ShouldPanic)
	})
}

func TestString_Set(t *testing.T) {
	C.Convey("String.Set", t, func() {
		val := optional.String{}
		C.So(val.Exists(), C.ShouldBeFalse)
		val.Set("ember")
		C.So(val.Exists(), C.ShouldBeTrue)
		C.So(val.Get(), C.ShouldEqual, "ember")
	})
}

func TestString_UnmarshalJSON(t *testing.T) {
	C.Convey("String.UnmarshalJSON", t, func() {
		val := optional.String{}
		C.So(val.UnmarshalJSON([]byte(`"tungulus"`)), C.ShouldBeNil)
		C.So(val.Get(), C.ShouldEqual, "tungulus")
	})
}

func TestString_MarshalJSON(t *testing.T) {
	C.Convey("String.MarshalJSON", t, func() {
		val := optional.NewString("phobia")
		a, b := val.MarshalJSON()
		C.So(b, C.ShouldBeNil)
		C.So(string(a), C.ShouldEqual, `"phobia"`)
	})

}
