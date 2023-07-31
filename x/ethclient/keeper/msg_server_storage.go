package keeper

import (
	"context"

	"ethclient/x/ethclient/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateStorage(goCtx context.Context, msg *types.MsgCreateStorage) (*types.MsgCreateStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetStorage(ctx, msg.Address, msg.Storage, msg.Block)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	var storage = types.Storage{
		Address: msg.Address,
		Storage: msg.Storage,
		Block:   msg.Block,
		Value:   "TODO",
		Args: &types.Args{
			Address: msg.Address,
			Storage: msg.Storage,
			Block:   msg.Block,
		},
	}

	k.SetStorage(
		ctx,
		storage,
		msg.Address,
		msg.Storage,
		msg.Block,
	)
	return &types.MsgCreateStorageResponse{}, nil
}
