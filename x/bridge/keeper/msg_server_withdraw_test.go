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

func TestWithdrawMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.BridgeKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateWithdraw{Creator: creator,
			TxHash: strconv.Itoa(i),
		}
		_, err := srv.CreateWithdraw(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetWithdraw(ctx,
			expected.TxHash,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestWithdrawMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateWithdraw
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateWithdraw{Creator: creator,
				TxHash: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateWithdraw{Creator: "B",
				TxHash: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateWithdraw{Creator: creator,
				TxHash: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BridgeKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateWithdraw{Creator: creator,
				TxHash: strconv.Itoa(0),
			}
			_, err := srv.CreateWithdraw(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateWithdraw(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetWithdraw(ctx,
					expected.TxHash,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestWithdrawMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteWithdraw
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteWithdraw{Creator: creator,
				TxHash: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteWithdraw{Creator: "B",
				TxHash: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteWithdraw{Creator: creator,
				TxHash: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BridgeKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateWithdraw(wctx, &types.MsgCreateWithdraw{Creator: creator,
				TxHash: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteWithdraw(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetWithdraw(ctx,
					tc.request.TxHash,
				)
				require.False(t, found)
			}
		})
	}
}
