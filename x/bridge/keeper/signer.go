package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/x/bridge/types"
)

// GetSignerCount get the total number of signer
func (k Keeper) GetSignerCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SignerCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSignerCount set the total number of signer
func (k Keeper) SetSignerCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SignerCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSigner appends a signer in the store with a new id and update the count
func (k Keeper) AppendSigner(
	ctx sdk.Context,
	signer types.Signer,
) uint64 {
	// Create the signer
	count := k.GetSignerCount(ctx)

	// Set the ID of the appended value
	signer.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignerKey))
	appendedValue := k.cdc.MustMarshal(&signer)
	store.Set(GetSignerIDBytes(signer.Id), appendedValue)

	// Update signer count
	k.SetSignerCount(ctx, count+1)

	return count
}

// SetSigner set a specific signer in the store
func (k Keeper) SetSigner(ctx sdk.Context, signer types.Signer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignerKey))
	b := k.cdc.MustMarshal(&signer)
	store.Set(GetSignerIDBytes(signer.Id), b)
}

// GetSigner returns a signer from its id
func (k Keeper) GetSigner(ctx sdk.Context, id uint64) (val types.Signer, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignerKey))
	b := store.Get(GetSignerIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSigner removes a signer from the store
func (k Keeper) RemoveSigner(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignerKey))
	store.Delete(GetSignerIDBytes(id))
}

// GetAllSigner returns all signer
func (k Keeper) GetAllSigner(ctx sdk.Context) (list []types.Signer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignerKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Signer
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSignerIDBytes returns the byte representation of the ID
func GetSignerIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSignerIDFromBytes returns ID in uint64 format from a byte array
func GetSignerIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
