package q567_checkInclusion

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_checkInclusion(t *testing.T) {
	Convey("Test_checkInclusion", t, func() {
		So(checkInclusion("ab", "eidbaooo"), ShouldBeTrue)
		So(checkInclusion("ab", "eidboaoo"), ShouldBeFalse)
	})

	Convey("Test_getDictn", t, func() {
		So(getDict("ab"), ShouldResemble, map[byte]int{'a': 1, 'b': 1})
	})
}
