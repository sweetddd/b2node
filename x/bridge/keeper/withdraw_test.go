package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/evmos/ethermint/testutil/bridge/keeper"
	"github.com/evmos/ethermint/testutil/bridge/nullify"
	"github.com/evmos/ethermint/x/bridge/keeper"
	"github.com/evmos/ethermint/x/bridge/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNWithdraw(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Withdraw {
	items := make([]types.Withdraw, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetWithdraw(ctx, items[i])
	}
	return items
}

func TestWithdrawGet(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNWithdraw(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetWithdraw(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestWithdrawRemove(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNWithdraw(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveWithdraw(ctx,
			item.Index,
		)
		_, found := keeper.GetWithdraw(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestWithdrawGetAll(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNWithdraw(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllWithdraw(ctx)),
	)
}
