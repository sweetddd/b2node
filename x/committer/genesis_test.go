package committer_test

import (
	"testing"

	"github.com/evmos/ethermint/testutil"
	keepertest "github.com/evmos/ethermint/testutil/keeper"

	"github.com/evmos/ethermint/x/committer"
	"github.com/evmos/ethermint/x/committer/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	addr := testutil.AccAddress()
	genesisState.Committers = types.Committer{
		CommitterList: []string{addr},
	}
	genesisState.Params.AdminPolicy = []*types.AdminPolicy{
		{
			PolicyType: types.PolicyType_group1,
			Address:    addr,
		},
	}

	k, ctx := keepertest.CommitterKeeper(t)

	committer.InitGenesis(ctx, *k, nil, genesisState)
	got := committer.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// nullify.Fill(&genesisState)
	// nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
