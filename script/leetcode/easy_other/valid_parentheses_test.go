package easyother

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	res := isValid("{}[]")
	assert.Equal(t, res, true)
}
