package keeper //nolint:dupl

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/x/bridge/types"
)

// SetRollupTx set a specific rollupTx in the store from its index
func (k Keeper) SetRollupTx(ctx sdk.Context, rollupTx types.RollupTx) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RollupTxKeyPrefix))
	b := k.cdc.MustMarshal(&rollupTx)
	store.Set(types.RollupTxKey(
		rollupTx.TxHash,
	), b)
}

// GetRollupTx returns a rollupTx from its index
func (k Keeper) GetRollupTx(
	ctx sdk.Context,
	txHash string,
) (val types.RollupTx, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RollupTxKeyPrefix))

	b := store.Get(types.RollupTxKey(
		txHash,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRollupTx removes a rollupTx from the store
func (k Keeper) RemoveRollupTx(
	ctx sdk.Context,
	txHash string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RollupTxKeyPrefix))
	store.Delete(types.RollupTxKey(
		txHash,
	))
}

// GetAllRollupTx returns all rollupTx
func (k Keeper) GetAllRollupTx(ctx sdk.Context) (list []types.RollupTx) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RollupTxKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RollupTx
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
