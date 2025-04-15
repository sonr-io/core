package keeper

import (
	"context"

	snrctx "github.com/sonr-io/core/ctx"
	"github.com/sonr-io/core/x/dwn/types"
)

var _ types.QueryServer = Querier{}

type Querier struct {
	Keeper
}

func NewQuerier(keeper Keeper) Querier {
	return Querier{Keeper: keeper}
}

func (k Querier) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := snrctx.FromGoContext(c)
	p, err := k.Keeper.Params.Get(ctx.SDKContext())
	if err != nil {
		return nil, err
	}
	return &types.QueryParamsResponse{Params: &p}, nil
}

// Spawn implements types.QueryServer.
func (k Querier) Spawn(goCtx context.Context, req *types.QuerySpawnRequest) (*types.QuerySpawnResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QuerySpawnResponse{}, nil
}

// Check implements types.QueryServer.
func (k Querier) Check(goCtx context.Context, req *types.QueryCheckRequest) (*types.QueryCheckResponse, error) {
	ctx := snrctx.FromGoContext(goCtx)
	return &types.QueryCheckResponse{
		Address: ctx.Peer().Addr.String(),
	}, nil
}
