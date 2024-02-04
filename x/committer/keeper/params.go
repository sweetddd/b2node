package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/x/committer/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	store := ctx.KVStore(k.storeKey)
	var params types.Params
	k.cdc.MustUnmarshal(store.Get(types.KeyPrefixParams), &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&params)

	store.Set(types.KeyPrefixParams, bz)
}
