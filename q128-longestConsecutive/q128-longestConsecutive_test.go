package q128_longestConsecutive

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_longestConsecutive(t *testing.T) {

	Convey("Test_longestConsecutive", t, func() {
		So(longestConsecutive([]int{100, 4, 200, 1, 3, 2}), ShouldEqual, 4)
		So(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}), ShouldEqual, 9)
	})
}
