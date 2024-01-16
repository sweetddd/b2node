package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
    

	"github.com/evmos/ethermint/x/committer/types"
)

type (
	Keeper struct {
		cdc      	codec.BinaryCodec
		storeKey 	storetypes.StoreKey
		memKey   	storetypes.StoreKey
		paramstore	paramtypes.Subspace
        
		
	}
)

func NewKeeper(
    cdc codec.BinaryCodec,
    storeKey,
    memKey storetypes.StoreKey,
    ps paramtypes.Subspace,
    
    
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
        cdc:      	cdc,
        storeKey: 	storeKey,
        memKey:   	memKey,
        paramstore:	ps,
        
		
	}
}



func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetLastProposalId(ctx sdk.Context) uint64 {
	// TODO: implement
	return 0
}

func (k Keeper) SetProposal(ctx sdk.Context, proposal types.Proposal) {
	// TODO: implement
}

func (k Keeper) GetProposal(ctx sdk.Context, id uint64) types.Proposal {
	// TODO: implement
	return types.Proposal{}
}

func (k Keeper) AddCommitter(ctx sdk.Context, committer types.Committer) {
	// TODO: implement
}

func (k Keeper) GetAllCommitters(ctx sdk.Context) []types.Committer {
	// TODO: implement
	return []types.Committer{}
}

func (k Keeper) IsExistCommitter(ctx sdk.Context, address string) bool {
	// TODO: implement
	return false
}
