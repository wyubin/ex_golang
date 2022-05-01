package easyother

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	res := strStr("aahaa", "haa")
	assert.Equal(t, 2, res)
}
