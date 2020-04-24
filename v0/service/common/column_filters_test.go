package common_test

import (
	"encoding/json"
	"errors"
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/common"
	C "github.com/smartystreets/goconvey/convey"
	"testing"
)

type nonMarshalable byte

func (n nonMarshalable) MarshalJSON() ([]byte, error) {
	return nil, errors.New("foo")
}

func TestColumnFilters_Append(t *testing.T) {
	C.Convey("ColumnFilters.Append", t, func() {
		C.Convey("Valid value", func() {
			test := common.ColumnFilters{}
			C.So(test.Append("foo", "bar", "fizz"), C.ShouldBeNil)
			C.So(test["foo"]["bar"], C.ShouldResemble, json.RawMessage(`"fizz"`))
		})

		C.Convey("Invalid value", func() {
			test := common.ColumnFilters{}
			C.So(test.Append("foo", "bar", nonMarshalable(0)), C.ShouldNotBeNil)
		})
	})
}
