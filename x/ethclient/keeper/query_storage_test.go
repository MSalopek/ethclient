package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "ethclient/testutil/keeper"
	"ethclient/testutil/nullify"
	"ethclient/x/ethclient/types"
)

func TestStorageQuery(t *testing.T) {
	keeper, ctx := keepertest.EthclientKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	addr := "mock-addr"
	key := "mock-key"
	block := int64(1)
	item := createTestStorage(keeper, ctx, addr, key, block)
	tests := []struct {
		desc     string
		request  *types.QueryGetStorageRequest
		response *types.QueryGetStorageResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetStorageRequest{
				Address: addr,
				Storage: key,
				Block:   block,
			},
			response: &types.QueryGetStorageResponse{Storage: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.QueryStorage(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
