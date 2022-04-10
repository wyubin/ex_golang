package main

import (
	"context"
)

func GetContextValue(ctx context.Context, k string) string {
	v, ok := ctx.Value(k).(string)
	if !ok {
		return ""
	}
	return v
}

func SetContextWithKeyVal(ctx context.Context, key, val string) context.Context {
	return context.WithValue(ctx, key, val)
}
