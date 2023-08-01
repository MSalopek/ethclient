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

	// Fetch ETH storage & proofs using the ETH RPC client
	s, err := k.EthCreateStorage(ctx, msg.Address, msg.Storage, msg.Block)
	if err != nil {
		return &types.MsgCreateStorageResponse{}, err
	}

	return &types.MsgCreateStorageResponse{
		Storage: &s,
	}, nil
}
