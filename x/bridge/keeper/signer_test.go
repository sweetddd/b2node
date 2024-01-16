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

func createNSigner(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Signer {
	items := make([]types.Signer, n)
	for i := range items {
		items[i].Id = keeper.AppendSigner(ctx, items[i])
	}
	return items
}

func TestSignerGet(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNSigner(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSigner(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestSignerRemove(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNSigner(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSigner(ctx, item.Id)
		_, found := keeper.GetSigner(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSignerGetAll(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNSigner(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSigner(ctx)),
	)
}

func TestSignerCount(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	items := createNSigner(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSignerCount(ctx))
}
