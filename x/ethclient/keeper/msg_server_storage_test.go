package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "ethclient/testutil/keeper"
	"ethclient/x/ethclient/keeper"
	"ethclient/x/ethclient/types"
)

func TestStorageMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.EthclientKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	expected := &types.MsgCreateStorage{Creator: creator}
	_, err := srv.CreateStorage(wctx, expected)
	require.NoError(t, err)
	rst, found := k.GetStorage(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}

func TestStorageMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateStorage
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateStorage{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateStorage{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.EthclientKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateStorage{Creator: creator}
			_, err := srv.CreateStorage(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateStorage(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetStorage(ctx)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestStorageMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteStorage
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteStorage{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteStorage{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.EthclientKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateStorage(wctx, &types.MsgCreateStorage{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteStorage(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetStorage(ctx)
				require.False(t, found)
			}
		})
	}
}
