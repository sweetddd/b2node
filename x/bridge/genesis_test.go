package bridge_test

import (
	"testing"

	keepertest "github.com/evmos/ethermint/testutil/bridge/keeper"
	"github.com/evmos/ethermint/testutil/bridge/nullify"
	"github.com/evmos/ethermint/x/bridge"
	"github.com/evmos/ethermint/x/bridge/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SignerGroupList: []types.SignerGroup{
			{
				Name: "0",
			},
			{
				Name: "1",
			},
		},
		CallerGroupList: []types.CallerGroup{
			{
				Name: "0",
			},
			{
				Name: "1",
			},
		},
		DepositList: []types.Deposit{
			{
				TxHash: "0",
			},
			{
				TxHash: "1",
			},
		},
		WithdrawList: []types.Withdraw{
			{
				TxHash: "0",
			},
			{
				TxHash: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BridgeKeeper(t)
	bridge.InitGenesis(ctx, *k, genesisState)
	got := bridge.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.SignerGroupList, got.SignerGroupList)
	require.ElementsMatch(t, genesisState.CallerGroupList, got.CallerGroupList)
	require.ElementsMatch(t, genesisState.DepositList, got.DepositList)
	require.ElementsMatch(t, genesisState.WithdrawList, got.WithdrawList)
	// this line is used by starport scaffolding # genesis/test/assert
}
