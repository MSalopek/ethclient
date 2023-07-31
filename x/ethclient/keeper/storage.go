package keeper

import (
	"ethclient/x/ethclient/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetStorage set storage in the store
func (k Keeper) SetStorage(ctx sdk.Context, storage types.Storage, address, key, block string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetStorageKey(address, key, block))
	b := k.cdc.MustMarshal(&storage)
	store.Set([]byte{0}, b)
}

// GetStorage returns storage
func (k Keeper) GetStorage(ctx sdk.Context, address, key, block string) (val types.Storage, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetStorageKey(address, key, block))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
