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

func TestSignerGroupMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.BridgeKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateSignerGroup{Creator: creator,
			Name: strconv.Itoa(i),
		}
		_, err := srv.CreateSignerGroup(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetSignerGroup(ctx,
			expected.Name,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestSignerGroupMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateSignerGroup
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateSignerGroup{Creator: creator,
				Name: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateSignerGroup{Creator: "B",
				Name: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateSignerGroup{Creator: creator,
				Name: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BridgeKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateSignerGroup{Creator: creator,
				Name: strconv.Itoa(0),
			}
			_, err := srv.CreateSignerGroup(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateSignerGroup(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetSignerGroup(ctx,
					expected.Name,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestSignerGroupMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteSignerGroup
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteSignerGroup{Creator: creator,
				Name: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteSignerGroup{Creator: "B",
				Name: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteSignerGroup{Creator: creator,
				Name: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BridgeKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateSignerGroup(wctx, &types.MsgCreateSignerGroup{Creator: creator,
				Name: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteSignerGroup(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetSignerGroup(ctx,
					tc.request.Name,
				)
				require.False(t, found)
			}
		})
	}
}
