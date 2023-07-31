package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/status"

	"ethclient/testutil/network"
	"ethclient/testutil/nullify"
	"ethclient/x/ethclient/client/cli"
	"ethclient/x/ethclient/types"
)

func networkWithStorageObjects(t *testing.T) (*network.Network, types.Storage) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	storage := &types.Storage{}
	nullify.Fill(&storage)
	state.Storage = storage
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), *state.Storage
}

func TestShowStorage(t *testing.T) {
	net, obj := networkWithStorageObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	tests := []struct {
		desc string
		args []string
		err  error
		obj  types.Storage
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowStorage(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetStorageResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.Storage)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.Storage),
				)
			}
		})
	}
}
