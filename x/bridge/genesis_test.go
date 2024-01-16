package bridge_test

import (
	keepertest "github.com/evmos/ethermint/testutil/bridge/keeper"
	"github.com/evmos/ethermint/testutil/bridge/nullify"
	"testing"

	"github.com/evmos/ethermint/x/bridge"
	"github.com/evmos/ethermint/x/bridge/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		CallerList: []types.Caller{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		CallerCount: 2,
		DepositList: []types.Deposit{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		WithdrawList: []types.Withdraw{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		SignerList: []types.Signer{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		SignerCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BridgeKeeper(t)
	bridge.InitGenesis(ctx, *k, genesisState)
	got := bridge.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.CallerList, got.CallerList)
	require.Equal(t, genesisState.CallerCount, got.CallerCount)
	require.ElementsMatch(t, genesisState.DepositList, got.DepositList)
	require.ElementsMatch(t, genesisState.WithdrawList, got.WithdrawList)
	require.ElementsMatch(t, genesisState.SignerList, got.SignerList)
	require.Equal(t, genesisState.SignerCount, got.SignerCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
