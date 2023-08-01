package keeper

import (
	"ethclient/x/ethclient/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetStorage set storage in the store
func (k Keeper) SetStorage(ctx sdk.Context, storage types.Storage, address, key string, block int64) {
	store := ctx.KVStore(k.storeKey)
	storeKey := types.GetStorageKey(address, key, block)
	b := k.cdc.MustMarshal(&storage)
	store.Set(storeKey, b)
}

// GetStorage returns storage
func (k Keeper) GetStorage(ctx sdk.Context, address, key string, block int64) (val types.Storage, found bool) {
	store := ctx.KVStore(k.storeKey)
	storeKey := types.GetStorageKey(address, key, block)

	b := store.Get(storeKey)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetProof returns storage proof
func (k Keeper) GetProof(ctx sdk.Context, address, key string, block int64) (*types.StorageProof, bool) {
	store := ctx.KVStore(k.storeKey)
	storeKey := types.GetProofKey(address, key, block)

	b := store.Get(storeKey)
	if b == nil {
		return &types.StorageProof{}, false
	}

	var proof types.StorageProof
	k.cdc.MustUnmarshal(b, &proof)
	return &proof, true
}

// SetProof sets storage proof
func (k Keeper) SetProof(ctx sdk.Context, address, key string, block int64, proof *types.StorageProof) {
	store := ctx.KVStore(k.storeKey)
	storeKey := types.GetProofKey(address, key, block)
	b := k.cdc.MustMarshal(proof)
	store.Set(storeKey, b)
}
