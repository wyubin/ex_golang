package easyother

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	res := climbStairs(4)
	assert.Equal(t, 5, res)
}
