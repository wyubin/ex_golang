package easyother

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	n := []int{1, 2, 3, 0}
	merge(n, 3, []int{4}, 1)
	assert.Equal(t, []int{1, 2, 3, 4}, n)
}
