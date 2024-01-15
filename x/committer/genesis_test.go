package committer_test

import (
	"testing"

	keepertest "github.com/evmos/ethermint/testutil/keeper"
	//"github.com/evmos/ethermint/testutil/nullify"
	"github.com/evmos/ethermint/x/committer"
	"github.com/evmos/ethermint/x/committer/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CommitterKeeper(t)
	committer.InitGenesis(ctx, *k, genesisState)
	got := committer.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	//nullify.Fill(&genesisState)
	//nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
