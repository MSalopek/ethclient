package keeper_test

import (
	"testing"

	testkeeper "ethclient/testutil/keeper"
	"ethclient/x/ethclient/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EthclientKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
