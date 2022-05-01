package easyother

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	res := romanToInt("LIX")
	assert.Equal(t, res, 59)
}
