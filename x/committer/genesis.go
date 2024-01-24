package committer

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/x/committer/keeper"
	"github.com/evmos/ethermint/x/committer/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, ak types.AccountKeeper, genState types.GenesisState) {
  // this line is used by starport scaffolding # genesis/module/init

	if genState.Params.TimeoutBlocks == 0 {
		genState.Params.TimeoutBlocks = types.DefaultTimeoutBlockPeriod
	}

	if len(genState.Params.AdminPolicy) == 0 {
		// if no admin policy is set, set the first account as admin
		// TODO: maybe we should use a more secure way to set admin policy
		accs := ak.GetAllAccounts(ctx)
		genState.Params.AdminPolicy = []*types.AdminPolicy{
			{
				Address: accs[0].GetAddress().String(),
				PolicyType:  types.PolicyType_group1,
			},
		}
	}

	k.SetParams(ctx, genState.Params)
	if len(genState.Committers.CommitterList) > 0{ 
		k.SetCommitter(ctx, genState.Committers)
	} else {
		// set committer list from all accounts
		// TODO: maybe we should use a more secure way to set committer list
		accs := ak.GetAllAccounts(ctx)
		var committers []string
		for _, acc := range accs {
			committers = append(committers, acc.GetAddress().String())
		}
		k.SetCommitter(ctx, types.Committer{CommitterList: committers})
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.Committers = k.GetAllCommitters(ctx)

    // this line is used by starport scaffolding # genesis/module/export

    return genesis
}
