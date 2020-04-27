package param_test

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/VEuPathDB/lib-go-wdk-api/v0/model/param"
)

func TestNumberRange_MarshalJSON(t *testing.T) {
	Convey("Marshaling a NumberRange param to JSON", t, func() {
		Convey("When param type is NumberRange", func() {
			Convey("Should succeed", func() {
				_, err := json.Marshal(param.NumberRange{Base: param.Base{Kind: param.KindNumberRange}})
				So(err, ShouldBeNil)
			})
		})

		Convey("When param type is not NumberRange", func() {
			Convey("Should fail", func() {
				tests := []param.Kind{
					param.KindDate,
					param.KindDateRange,
					param.KindFilter,
					param.KindDataset,
					param.KindAnswer,
					param.KindNumber,
					param.KindString,
					param.KindTimestamp,
					param.KindSingleVocab,
					param.KindMultiVocab,
				}

				for _, test := range tests {
					par := param.NumberRange{Base: param.Base{Kind: test}}
					_, err := json.Marshal(par)
					So(err, ShouldNotBeNil)
				}
			})
		})
	})
}
