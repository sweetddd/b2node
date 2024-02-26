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

func createNRollupTx(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.RollupTx {
	items := make([]types.RollupTx, n)
	for i := range items {
		items[i].TxHash = strconv.Itoa(i)

		keeper.SetRollupTx(ctx, items[i])
	}
	return items
}

func TestRollupTxGet(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNRollupTx(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRollupTx(ctx,
			item.TxHash,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestRollupTxRemove(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNRollupTx(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRollupTx(ctx,
			item.TxHash,
		)
		_, found := keeper.GetRollupTx(ctx,
			item.TxHash,
		)
		require.False(t, found)
	}
}

func TestRollupTxGetAll(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNRollupTx(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRollupTx(ctx)),
	)
}
