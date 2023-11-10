package bitcoinindexer_test

import (
	"testing"

	keepertest "github.com/evmos/ethermint/testutil/keeper"
	// "github.com/evmos/ethermint/testutil/nullify"
	"github.com/evmos/ethermint/x/bitcoinindexer"
	"github.com/evmos/ethermint/x/bitcoinindexer/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BitcoinindexerKeeper(t)
	bitcoinindexer.InitGenesis(ctx, *k, genesisState)
	got := bitcoinindexer.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// nullify.Fill(&genesisState)
	// nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
