package internal

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/peer"
)

type SonrContext interface {
	SDKContext() sdk.Context
	Peer() *peer.Peer
	IPFSRedirectURL(cid string) string
}

// Context is a wrapper around the Cosmos SDK context that provides additional
// functionality.
type sonrContext struct {
	sdk  sdk.Context
	peer *peer.Peer
}

func (c sonrContext) SDKContext() sdk.Context {
	return c.sdk
}

func (c sonrContext) Peer() *peer.Peer {
	return c.peer
}

func (c sonrContext) IPFSRedirectURL(cid string) string {
	return "http://ipfs.sonr.land/ipfs/" + cid
}

// NewContext creates a new context with the given Cosmos SDK context and peer
// information.
func NewContext(ctx sdk.Context, p *peer.Peer) SonrContext {
	return &sonrContext{
		sdk:  ctx,
		peer: p,
	}
}

// FromSDKContext returns the context from the given context.
func FromSDKContext(ctx sdk.Context) SonrContext {
	p, _ := peer.FromContext(ctx)
	return &sonrContext{
		sdk:  ctx,
		peer: p,
	}
}

// FromGoContext returns the context from the given context.
func FromGoContext(ctx context.Context) SonrContext {
	sctx := sdk.UnwrapSDKContext(ctx)
	p, _ := peer.FromContext(ctx)
	return &sonrContext{
		sdk:  sctx,
		peer: p,
	}
}
