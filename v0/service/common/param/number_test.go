package param_test

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/common/param"
)

func TestNumber_MarshalJSON(t *testing.T) {
	Convey("Marshaling a Number param to JSON", t, func() {
		Convey("When param type is Number", func() {
			Convey("Should succeed", func() {
				_, err := json.Marshal(param.Number{Base: param.Base{Kind: param.KindNumber}})
				So(err, ShouldBeNil)
			})
		})

		Convey("When param type is not Number", func() {
			Convey("Should fail", func() {
				tests := []param.Kind{
					param.KindDate,
					param.KindDateRange,
					param.KindFilter,
					param.KindDataset,
					param.KindAnswer,
					param.KindNumberRange,
					param.KindString,
					param.KindTimestamp,
					param.KindSingleVocab,
					param.KindMultiVocab,
				}

				for _, test := range tests {
					param := param.Number{Base: param.Base{Kind: test}}
					_, err := json.Marshal(param)
					So(err, ShouldNotBeNil)
				}
			})
		})
	})
}
