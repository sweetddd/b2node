package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/x/bridge/types"
)

// SetWithdraw set a specific withdraw in the store from its index
func (k Keeper) SetWithdraw(ctx sdk.Context, withdraw types.Withdraw) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawKeyPrefix))
	b := k.cdc.MustMarshal(&withdraw)
	store.Set(types.WithdrawKey(
		withdraw.TxId,
	), b)
}

// GetWithdraw returns a withdraw from its index
func (k Keeper) GetWithdraw(
	ctx sdk.Context,
	txID string,
) (val types.Withdraw, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawKeyPrefix))

	b := store.Get(types.WithdrawKey(
		txID,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveWithdraw removes a withdraw from the store
func (k Keeper) RemoveWithdraw(
	ctx sdk.Context,
	txID string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawKeyPrefix))
	store.Delete(types.WithdrawKey(
		txID,
	))
}

// GetAllWithdraw returns all withdraw
func (k Keeper) GetAllWithdraw(ctx sdk.Context) (list []types.Withdraw) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Withdraw
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetStatusIndex use status and txHash composite key
func (k Keeper) SetStatusIndex(ctx sdk.Context, status string, txID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawStatusKeyPrefix))
	store.Set(types.WithdrawStatusKey(
		status,
		txID,
	), []byte(txID))
}

// RemoveStatusIndex removes a status index
func (k Keeper) RemoveStatusIndex(
	ctx sdk.Context,
	status string,
	txID string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawStatusKeyPrefix))
	store.Delete(types.WithdrawStatusKey(
		status,
		txID,
	))
}

// GetAllWithdrawByStatus returns all withdraw which have the same status
func (k Keeper) GetAllWithdrawByStatus(ctx sdk.Context, status string) (list []types.Withdraw) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawStatusKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte(status))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		txID := string(iterator.Value())
		val, found := k.GetWithdraw(ctx, txID)
		if !found {
			continue
		}
		list = append(list, val)
	}

	return
}
