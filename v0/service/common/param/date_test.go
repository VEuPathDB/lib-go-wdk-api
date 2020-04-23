package param_test

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/common/param"
)

func TestDate_MarshalJSON(t *testing.T) {
	Convey("Marshaling a Date param to JSON", t, func() {
		Convey("When param type is Date", func() {
			Convey("Should succeed", func() {
				_, err := json.Marshal(param.Date{Base: param.Base{Kind: param.KindDate}})
				So(err, ShouldBeNil)
			})
		})

		Convey("When param type is not Date", func() {
			Convey("Should fail", func() {
				tests := []param.Kind{
					param.KindDateRange,
					param.KindFilter,
					param.KindDataset,
					param.KindAnswer,
					param.KindNumber,
					param.KindNumberRange,
					param.KindString,
					param.KindTimestamp,
					param.KindSingleVocab,
					param.KindMultiVocab,
				}

				for _, test := range tests {
					param := param.Date{Base: param.Base{Kind: test}}
					_, err := json.Marshal(param)
					So(err, ShouldNotBeNil)
				}
			})
		})
	})
}
