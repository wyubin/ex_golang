package easyother

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTitleToNumber(t *testing.T) {
	res := titleToNumber("AC")
	assert.Equal(t, res, 29)
}
