package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/evmos/ethermint/testutil/bridge/keeper"
	"github.com/evmos/ethermint/x/bridge/keeper"
	"github.com/evmos/ethermint/x/bridge/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestDepositMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.BridgeKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	expected := &types.MsgCreateDeposit{Creator: creator,
		TxHash: "expected error",
	}
	_, err := srv.CreateDeposit(wctx, expected)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
	k.SetParams(ctx, types.DefaultParams())
	srv.CreateCallerGroup(wctx, &types.MsgCreateCallerGroup{
		Creator: creator,
		Name:    "caller group",
		Admin:   creator,
		Members: []string{creator},
	})
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateDeposit{Creator: creator,
			TxHash: strconv.Itoa(i),
		}
		_, err := srv.CreateDeposit(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetDeposit(ctx,
			expected.TxHash,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestDepositMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateDeposit
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateDeposit{Creator: creator,
				TxHash: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateDeposit{Creator: "B",
				TxHash: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateDeposit{Creator: creator,
				TxHash: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BridgeKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			k.SetParams(ctx, types.DefaultParams())
			srv.CreateCallerGroup(wctx, &types.MsgCreateCallerGroup{
				Creator: creator,
				Name:    "caller group",
				Admin:   creator,
				Members: []string{creator},
			})
			expected := &types.MsgCreateDeposit{Creator: creator,
				TxHash: strconv.Itoa(0),
			}
			_, err := srv.CreateDeposit(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateDeposit(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetDeposit(ctx,
					expected.TxHash,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestDepositMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteDeposit
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteDeposit{Creator: creator,
				TxHash: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteDeposit{Creator: "B",
				TxHash: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteDeposit{Creator: creator,
				TxHash: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BridgeKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			k.SetParams(ctx, types.DefaultParams())
			srv.CreateCallerGroup(wctx, &types.MsgCreateCallerGroup{
				Creator: creator,
				Name:    "caller group",
				Admin:   creator,
				Members: []string{creator},
			})
			_, err := srv.CreateDeposit(wctx, &types.MsgCreateDeposit{Creator: creator,
				TxHash: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteDeposit(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetDeposit(ctx,
					tc.request.TxHash,
				)
				require.False(t, found)
			}
		})
	}
}
