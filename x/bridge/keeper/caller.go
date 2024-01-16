package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/x/bridge/types"
)

// GetCallerCount get the total number of caller
func (k Keeper) GetCallerCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CallerCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetCallerCount set the total number of caller
func (k Keeper) SetCallerCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CallerCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendCaller appends a caller in the store with a new id and update the count
func (k Keeper) AppendCaller(
	ctx sdk.Context,
	caller types.Caller,
) uint64 {
	// Create the caller
	count := k.GetCallerCount(ctx)

	// Set the ID of the appended value
	caller.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CallerKey))
	appendedValue := k.cdc.MustMarshal(&caller)
	store.Set(GetCallerIDBytes(caller.Id), appendedValue)

	// Update caller count
	k.SetCallerCount(ctx, count+1)

	return count
}

// SetCaller set a specific caller in the store
func (k Keeper) SetCaller(ctx sdk.Context, caller types.Caller) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CallerKey))
	b := k.cdc.MustMarshal(&caller)
	store.Set(GetCallerIDBytes(caller.Id), b)
}

// GetCaller returns a caller from its id
func (k Keeper) GetCaller(ctx sdk.Context, id uint64) (val types.Caller, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CallerKey))
	b := store.Get(GetCallerIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCaller removes a caller from the store
func (k Keeper) RemoveCaller(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CallerKey))
	store.Delete(GetCallerIDBytes(id))
}

// GetAllCaller returns all caller
func (k Keeper) GetAllCaller(ctx sdk.Context) (list []types.Caller) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CallerKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Caller
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCallerIDBytes returns the byte representation of the ID
func GetCallerIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetCallerIDFromBytes returns ID in uint64 format from a byte array
func GetCallerIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
