package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "ethclient/testutil/keeper"
	"ethclient/testutil/nullify"
	"ethclient/x/ethclient/keeper"
	"ethclient/x/ethclient/types"
)

func createTestStorage(keeper *keeper.Keeper, ctx sdk.Context, addr, key string, block int64) types.Storage {
	item := types.Storage{}
	keeper.SetStorage(ctx, item, addr, key, block)
	return item
}

func TestStorageCRUD(t *testing.T) {
	keeper, ctx := keepertest.EthclientKeeper(t)
	addr := "mock-addr"
	key := "mock-key"
	block := int64(0)
	item := createTestStorage(keeper, ctx, addr, key, block)
	rst, found := keeper.GetStorage(ctx, addr, key, block)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}
