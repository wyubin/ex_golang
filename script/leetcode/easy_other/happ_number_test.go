package easyother

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsHappy(t *testing.T) {
	// res := sumDigitSquare(11)
	// assert.Equal(t, 2, res)
	res := isHappy(19)
	assert.Equal(t, true, res)

}
