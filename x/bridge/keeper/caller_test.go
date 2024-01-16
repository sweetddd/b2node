package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/evmos/ethermint/testutil/bridge/keeper"
	"github.com/evmos/ethermint/testutil/bridge/nullify"
	"github.com/evmos/ethermint/x/bridge/keeper"
	"github.com/evmos/ethermint/x/bridge/types"
	"github.com/stretchr/testify/require"
)

func createNCaller(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Caller {
	items := make([]types.Caller, n)
	for i := range items {
		items[i].Id = keeper.AppendCaller(ctx, items[i])
	}
	return items
}

func TestCallerGet(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNCaller(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetCaller(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestCallerRemove(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNCaller(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCaller(ctx, item.Id)
		_, found := keeper.GetCaller(ctx, item.Id)
		require.False(t, found)
	}
}

func TestCallerGetAll(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNCaller(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCaller(ctx)),
	)
}

func TestCallerCount(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNCaller(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetCallerCount(ctx))
}
