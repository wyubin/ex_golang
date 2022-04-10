package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContentValue(t *testing.T) {
	ctx := context.Background()
	ctx1 := SetContextWithKeyVal(ctx, "key1", "val1")
	ctx1to2 := SetContextWithKeyVal(ctx1, "key2", "val2")
	ctx3 := SetContextWithKeyVal(ctx, "key3", "val3")
	// ctx1 不會有 key2
	assert.Equal(t, nil, ctx1.Value("key2"))
	// ctx1to2 會有 key1 跟 key2
	assert.Equal(t, "val1", ctx1to2.Value("key1"))
	assert.Equal(t, "val2", ctx1to2.Value("key2"))
	// ctx3 不會有 key1,只有 key3
	assert.Equal(t, nil, ctx3.Value("key1"))
	assert.Equal(t, "val3", ctx3.Value("key3"))
}
