package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/x/bridge/types"
)

// SetSignerGroup set a specific signerGroup in the store from its index
func (k Keeper) SetSignerGroup(ctx sdk.Context, signerGroup types.SignerGroup) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignerGroupKeyPrefix))
	b := k.cdc.MustMarshal(&signerGroup)
	store.Set(types.SignerGroupKey(
		signerGroup.Name,
	), b)
}

// GetSignerGroup returns a signerGroup from its index
func (k Keeper) GetSignerGroup(
	ctx sdk.Context,
	name string,
) (val types.SignerGroup, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignerGroupKeyPrefix))

	b := store.Get(types.SignerGroupKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetSignerGroupMembers(
	ctx sdk.Context,
	name string,
) []string {
	group, found := k.GetSignerGroup(ctx, name)
	if !found {
		return []string{}
	}
	return group.GetMembers()
}

func (k Keeper) GetSignerGroupThreshold(
	ctx sdk.Context,
	name string,
) uint32 {
	group, found := k.GetSignerGroup(ctx, name)
	if !found {
		return 0
	}
	return group.GetThreshold()
}

func (k Keeper) IsMemberInSignerGroup(
	ctx sdk.Context,
	name string,
	member string,
) bool {
	members := k.GetSignerGroupMembers(ctx, name)
	for _, v := range members {
		if v == member {
			return true
		}
	}
	return false
}

// RemoveSignerGroup removes a signerGroup from the store
func (k Keeper) RemoveSignerGroup(
	ctx sdk.Context,
	name string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignerGroupKeyPrefix))
	store.Delete(types.SignerGroupKey(
		name,
	))
}

// GetAllSignerGroup returns all signerGroup
func (k Keeper) GetAllSignerGroup(ctx sdk.Context) (list []types.SignerGroup) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SignerGroupKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SignerGroup
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
