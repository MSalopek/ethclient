package keeper

import (
	"context"

	"ethclient/x/ethclient/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Look up a storage value for a given address, storage key, and block number and return it if found.
func (k Keeper) QueryStorage(goCtx context.Context, req *types.QueryGetStorageRequest) (*types.QueryGetStorageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetStorage(ctx, req.Address, req.Key, req.Block)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	// TODO:
	return &types.QueryGetStorageResponse{Storage: val}, nil
}
