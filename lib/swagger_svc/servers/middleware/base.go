package middleware

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// aggregate middlewares to chain
func AggregateMW(mws ...endpoint.Middleware) endpoint.Middleware {
	if len(mws) > 0 {
		return endpoint.Chain(mws[0], mws[1:]...)
	}
	return mwChainEmpty()
}

// no mw apply
func mwChainEmpty() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			return next(ctx, request)
		}
	}
}
