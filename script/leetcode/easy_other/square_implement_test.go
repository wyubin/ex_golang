package easyother

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqrt(t *testing.T) {
	res := mySqrt(8)
	assert.Equal(t, 2, res)
}
