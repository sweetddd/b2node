package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/x/bridge/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.CallerGroupName(ctx),
		k.SignerGroupName(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// CallerGroupName returns the CallerGroupName param
func (k Keeper) CallerGroupName(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyCallerGroupName, &res)
	return
}

// SignerGroupName returns the SignerGroupName param
func (k Keeper) SignerGroupName(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeySignerGroupName, &res)
	return
}
