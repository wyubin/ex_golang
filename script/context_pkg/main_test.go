package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContentValue(t *testing.T) {
	val := "test"
	ctx := context.Background()
	ctx2 := SetContextWithKeyVal(ctx, "TestKey", val)
	// ctx 不會有 TestKey
	assert.NotEqual(t, val, GetContextValue(ctx, "TestKey"))
	// ctx2 才會有 TestKey
	assert.Equal(t, val, GetContextValue(ctx2, "TestKey"))

}
