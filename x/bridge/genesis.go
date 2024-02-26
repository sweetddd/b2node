package bridge

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/x/bridge/keeper"
	"github.com/evmos/ethermint/x/bridge/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, ak types.AccountKeeper, genState types.GenesisState) {
	if len(genState.SignerGroupList) == 0 || len(genState.CallerGroupList) == 0 {
		accs := ak.GetAllAccounts(ctx)
		adminAddress := ""
		for _, acc := range accs {
			if acc.GetAccountNumber() == 0 {
				adminAddress = acc.GetAddress().String()
			}
		}
		if len(genState.SignerGroupList) == 0 {
			genState.SignerGroupList = []types.SignerGroup{
				{
					Name:      "signer group",
					Admin:     adminAddress,
					Members:   []string{adminAddress},
					Threshold: 1,
					Creator:   adminAddress,
				},
			}
		}
		if len(genState.CallerGroupList) == 0 {
			genState.CallerGroupList = []types.CallerGroup{
				{
					Name:    "caller group",
					Admin:   adminAddress,
					Members: []string{adminAddress},
					Creator: adminAddress,
				},
			}
		}
	}
	// Set all the signerGroup
	for _, elem := range genState.SignerGroupList {
		k.SetSignerGroup(ctx, elem)
	}
	// Set all the callerGroup
	for _, elem := range genState.CallerGroupList {
		k.SetCallerGroup(ctx, elem)
	}
	// Set all the deposit
	for _, elem := range genState.DepositList {
		k.SetDeposit(ctx, elem)
	}
	// Set all the withdraw
	for _, elem := range genState.WithdrawList {
		k.SetWithdraw(ctx, elem)
	}
	// Set all the rollupTx
	for _, elem := range genState.RollupTxList {
		k.SetRollupTx(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.SignerGroupList = k.GetAllSignerGroup(ctx)
	genesis.CallerGroupList = k.GetAllCallerGroup(ctx)
	genesis.DepositList = k.GetAllDeposit(ctx)
	genesis.WithdrawList = k.GetAllWithdraw(ctx)
	genesis.RollupTxList = k.GetAllRollupTx(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
