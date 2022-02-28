package e

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type errorCustom struct {
	ApiError
}

type errorCustom1 struct {
	ApiError
}

func TestError(t *testing.T) {
	t.Run("Test one type error", oneErrorType)
	t.Run("Test two type error", twoErrorType)
}

// test same error type
func oneErrorType(t *testing.T) {
	// same error type with same Msg
	errCustom := errorCustom{}
	errCustom1 := fmt.Errorf("errCustom1: %w", errorCustom{})
	assert.True(t, errors.Is(errCustom1, errCustom))
	// same error type without same Msg
	errCustom.Msg = "revise Error"
	assert.False(t, errors.Is(errCustom1, errCustom))
	// same error source
	errCustom2 := fmt.Errorf("%w: errCustom2", errCustom1)
	errCustom3 := fmt.Errorf("%w: errCustom3", errorCustom{})
	assert.True(t, errors.Is(errCustom2, errorCustom{}))
	assert.True(t, errors.Is(errCustom3, errorCustom{}))
	t.Logf("errCustom2: %+v\n", errCustom2)
	t.Logf("errCustom3: %+v\n", errCustom3)
}

// test two error
func twoErrorType(t *testing.T) {
	errCustom := errorCustom{}
	errAnother := errorCustom1{}
	assert.False(t, errors.Is(errAnother, errCustom))
}
