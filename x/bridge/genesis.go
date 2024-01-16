package bridge

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/x/bridge/keeper"
	"github.com/evmos/ethermint/x/bridge/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the caller
	for _, elem := range genState.CallerList {
		k.SetCaller(ctx, elem)
	}

	// Set caller count
	k.SetCallerCount(ctx, genState.CallerCount)
	// Set all the deposit
	for _, elem := range genState.DepositList {
		k.SetDeposit(ctx, elem)
	}
	// Set all the withdraw
	for _, elem := range genState.WithdrawList {
		k.SetWithdraw(ctx, elem)
	}
	// Set all the signer
	for _, elem := range genState.SignerList {
		k.SetSigner(ctx, elem)
	}

	// Set signer count
	k.SetSignerCount(ctx, genState.SignerCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.CallerList = k.GetAllCaller(ctx)
	genesis.CallerCount = k.GetCallerCount(ctx)
	genesis.DepositList = k.GetAllDeposit(ctx)
	genesis.WithdrawList = k.GetAllWithdraw(ctx)
	genesis.SignerList = k.GetAllSigner(ctx)
	genesis.SignerCount = k.GetSignerCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
