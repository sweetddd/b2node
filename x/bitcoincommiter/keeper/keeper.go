package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/evmos/ethermint/x/bitcoincommiter/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
		config     *BITCOINConfig
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	homePath string,
) *Keeper {
	bitcoinConfig, _ := DefaultBITCOINConfig(homePath)
	return &Keeper{
		cdc:        cdc,
		memKey:     memKey,
		paramstore: ps,
		config:     bitcoinConfig,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
