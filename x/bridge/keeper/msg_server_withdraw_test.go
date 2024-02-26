package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
	expected := &types.MsgCreateWithdraw{Creator: creator,
		TxId: "expected error",
	}
	_, err := srv.CreateWithdraw(wctx, expected)
	require.ErrorIs(t, err, types.ErrNotCallerGroupMembers)
	k.SetParams(ctx, types.DefaultParams())
	srv.CreateCallerGroup(wctx, &types.MsgCreateCallerGroup{
		Creator: creator,
		Name:    "caller group",
		Admin:   creator,
		Members: []string{creator},
	})
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateWithdraw{Creator: creator,
			TxId: strconv.Itoa(i),
		}
		_, err := srv.CreateWithdraw(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetWithdraw(ctx,
			expected.TxId,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
	list := k.GetAllWithdrawByStatus(ctx, types.WithdrawStatus_WITHDRAW_STATUS_PENDING.String())
	require.Equal(t, 5, len(list))
}

func TestWithdrawMsgServerUpdate(t *testing.T) {
	creator := "A"
	signers := []string{"B", "C", "D"}

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateWithdraw
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateWithdraw{Creator: creator,
				TxId: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateWithdraw{Creator: "B",
				TxId: strconv.Itoa(0),
			},
			err: types.ErrNotCallerGroupMembers,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateWithdraw{Creator: creator,
				TxId: strconv.Itoa(100000),
			},
			err: types.ErrIndexNotExist,
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
			srv.CreateSignerGroup(wctx, &types.MsgCreateSignerGroup{
				Creator:   creator,
				Name:      "signer group",
				Admin:     creator,
				Threshold: 3,
				Members:   signers,
			})
			expected := &types.MsgCreateWithdraw{Creator: creator,
				TxId: strconv.Itoa(0),
			}
			_, err := srv.CreateWithdraw(wctx, expected)
			require.NoError(t, err)
			for _, signer := range signers {
				_, err := srv.SignWithdraw(wctx, &types.MsgSignWithdraw{Creator: signer, TxId: strconv.Itoa(0), Signature: signer})
				require.NoError(t, err)
			}
			_, err = srv.UpdateWithdraw(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetWithdraw(ctx,
					expected.TxId,
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
				TxId: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteWithdraw{Creator: "B",
				TxId: strconv.Itoa(0),
			},
			err: types.ErrNotOwner,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteWithdraw{Creator: creator,
				TxId: strconv.Itoa(100000),
			},
			err: types.ErrIndexNotExist,
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
			_, err := srv.CreateWithdraw(wctx, &types.MsgCreateWithdraw{Creator: creator,
				TxId: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteWithdraw(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetWithdraw(ctx,
					tc.request.TxId,
				)
				require.False(t, found)
			}
		})
	}
}
