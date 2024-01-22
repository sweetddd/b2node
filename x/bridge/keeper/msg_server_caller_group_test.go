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

func TestCallerGroupMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.BridgeKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateCallerGroup{Creator: creator,
			Name: strconv.Itoa(i),
		}
		_, err := srv.CreateCallerGroup(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetCallerGroup(ctx,
			expected.Name,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestCallerGroupMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateCallerGroup
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateCallerGroup{Creator: creator,
				Name: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateCallerGroup{Creator: "B",
				Name: strconv.Itoa(0),
			},
			err: types.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateCallerGroup{Creator: creator,
				Name: strconv.Itoa(100000),
			},
			err: types.ErrIndexNotExist,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BridgeKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateCallerGroup{Creator: creator,
				Name:  strconv.Itoa(0),
				Admin: creator,
			}
			_, err := srv.CreateCallerGroup(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateCallerGroup(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetCallerGroup(ctx,
					expected.Name,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestCallerGroupMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteCallerGroup
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteCallerGroup{Creator: creator,
				Name: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteCallerGroup{Creator: "B",
				Name: strconv.Itoa(0),
			},
			err: types.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteCallerGroup{Creator: creator,
				Name: strconv.Itoa(100000),
			},
			err: types.ErrIndexNotExist,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BridgeKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateCallerGroup(wctx, &types.MsgCreateCallerGroup{Creator: creator,
				Name:  strconv.Itoa(0),
				Admin: creator,
			})
			require.NoError(t, err)
			_, err = srv.DeleteCallerGroup(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetCallerGroup(ctx,
					tc.request.Name,
				)
				require.False(t, found)
			}
		})
	}
}
