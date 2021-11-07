package q233_countDigitOne

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getDigits(t *testing.T) {
	assert.Equal(t, []int{2, 3, 1}, getDigits(231))
}
