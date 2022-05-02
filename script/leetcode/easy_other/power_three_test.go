package easyother

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfThree(t *testing.T) {
	res := isPowerOfThree(27)
	assert.Equal(t, true, res)
}
