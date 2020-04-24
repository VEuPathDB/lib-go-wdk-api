package optional_test

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/optional"
	C "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewInt(t *testing.T) {
	C.Convey("NewInt", t, func() {
		C.So(optional.NewInt(3).Get(), C.ShouldEqual, 3)
		C.So(optional.NewInt(-7).Exists(), C.ShouldBeTrue)
	})
}

func TestInt_Clear(t *testing.T) {
	C.Convey("Int.Clear", t, func() {
		val := optional.NewInt(76)
		val.Clear()
		C.So(val.Exists(), C.ShouldBeFalse)
		C.So(func() {val.Get()}, C.ShouldPanic)
	})
}

func TestInt_Set(t *testing.T) {
	C.Convey("Int.Set", t, func() {
		val := optional.Int{}
		C.So(val.Exists(), C.ShouldBeFalse)
		val.Set(-86)
		C.So(val.Exists(), C.ShouldBeTrue)
		C.So(val.Get(), C.ShouldEqual, -86)
	})
}

func TestInt_UnmarshalJSON(t *testing.T) {
	C.Convey("Int.UnmarshalJSON", t, func() {
		val := optional.Int{}
		C.So(val.UnmarshalJSON([]byte(`-356`)), C.ShouldBeNil)
		C.So(val.Get(), C.ShouldEqual, -356)
	})
}

func TestInt_MarshalJSON(t *testing.T) {
	C.Convey("Int.MarshalJSON", t, func() {
		val := optional.NewInt(666)
		a, b := val.MarshalJSON()
		C.So(b, C.ShouldBeNil)
		C.So(string(a), C.ShouldEqual, `666`)
	})

}
