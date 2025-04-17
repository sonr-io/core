package keeper

import (
	"context"

	"cosmossdk.io/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/sonr-io/core/x/did/types"
)

type msgServer struct {
	k Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

// UpdateParams updates the x/did module parameters.
func (ms msgServer) UpdateParams(ctx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if ms.k.authority != msg.Authority {
		return nil, errors.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority; expected %s, got %s",
			ms.k.authority,
			msg.Authority,
		)
	}
	return nil, ms.k.Params.Set(ctx, msg.Params)
}

// ExecuteTx implements types.MsgServer.
func (ms msgServer) ExecuteTx(ctx context.Context, msg *types.MsgExecuteTx) (*types.MsgExecuteTxResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.MsgExecuteTxResponse{}, nil
}

// RegisterController implements types.MsgServer.
func (ms msgServer) RegisterController(ctx context.Context, msg *types.MsgRegisterController) (*types.MsgRegisterControllerResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("RegisterController is unimplemented")
	return &types.MsgRegisterControllerResponse{}, nil
}

// LinkVerificationMethod implements types.MsgServer.
func (ms msgServer) LinkVerificationMethod(ctx context.Context, msg *types.MsgLinkVerificationMethod) (*types.MsgLinkVerificationMethodResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("LinkVerificationMethod is unimplemented")
	return &types.MsgLinkVerificationMethodResponse{}, nil
}

// UnlinkVerificationMethod implements types.MsgServer.
func (ms msgServer) UnlinkVerificationMethod(ctx context.Context, msg *types.MsgUnlinkVerificationMethod) (*types.MsgUnlinkVerificationMethodResponse, error) {
	// ctx := sdk.UnwrapSDKContext(goCtx)
	panic("UnlinkVerificationMethod is unimplemented")
	return &types.MsgUnlinkVerificationMethodResponse{}, nil
}
