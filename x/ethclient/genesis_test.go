package ethclient_test

import (
	"testing"

	keepertest "ethclient/testutil/keeper"
	"ethclient/testutil/nullify"
	"ethclient/x/ethclient"
	"ethclient/x/ethclient/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EthclientKeeper(t)
	ethclient.InitGenesis(ctx, *k, genesisState)
	got := ethclient.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
