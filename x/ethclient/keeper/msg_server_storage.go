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
	_, isFound := k.GetStorage(ctx)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	var storage = types.Storage{
		Creator: msg.Creator,
	}

	k.SetStorage(
		ctx,
		storage,
	)
	return &types.MsgCreateStorageResponse{}, nil
}
