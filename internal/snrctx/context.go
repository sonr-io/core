package snrctx

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/peer"
)

// Context is a wrapper around the Cosmos SDK context that provides additional
// functionality.
type Context struct {
	sdk.Context
	*peer.Peer
}

// NewContext creates a new context with the given Cosmos SDK context and peer
// information.
func NewContext(ctx sdk.Context, p *peer.Peer) *Context {
	return &Context{
		Context: ctx,
		Peer:    p,
	}
}

// FromSDKContext returns the context from the given context.
func FromSDKContext(ctx sdk.Context) *Context {
	p, _ := peer.FromContext(ctx)
	return &Context{
		Context: ctx,
		Peer:    p,
	}
}

// FromGoContext returns the context from the given context.
func FromGoContext(ctx context.Context) *Context {
	sctx := sdk.UnwrapSDKContext(ctx)
	p, _ := peer.FromContext(ctx)
	return &Context{
		Context: sctx,
		Peer:    p,
	}
}
