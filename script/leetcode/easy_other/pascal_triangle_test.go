package easyother

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	res := generate(1)
	assert.Equal(t, [][]int{{1}}, res)
}
