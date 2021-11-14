package q423_originalDigits

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_originalDigits(t *testing.T) {

	Convey("Test_originalDigits", t, func() {
		So(originalDigits("owoztneoer"), ShouldEqual, "012")
	})

}
