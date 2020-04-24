package strategies_test

import (
	"github.com/VEuPathDB/lib-go-wdk-api/v0/service/users/strategies"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStep_HasPrimaryInput(t *testing.T) {
	convey.Convey("Step.HasPrimaryInput", t, func() {
		step := new(strategies.Step)
		convey.So(step.HasPrimaryInput(), convey.ShouldBeFalse)
		step.PrimaryInput = new(strategies.Step)
		convey.So(step.HasPrimaryInput(), convey.ShouldBeTrue)
	})
}

func TestStep_HasSecondaryInput(t *testing.T) {
	convey.Convey("Step.HasSecondaryInput", t, func() {
		step := new(strategies.Step)
		convey.So(step.HasSecondaryInput(), convey.ShouldBeFalse)
		step.SecondaryInput = new(strategies.Step)
		convey.So(step.HasSecondaryInput(), convey.ShouldBeTrue)
	})
}
