package validation_test

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/validation"
	C "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBundle_HasErrors(t *testing.T) {
	C.Convey("Bundle.HasErrors", t, func() {
		test := validation.Bundle{}
		C.So(test.HasErrors(), C.ShouldBeFalse)
		test.Errors = &validation.BundleErrors{}
		C.So(test.HasErrors(), C.ShouldBeTrue)
	})
}