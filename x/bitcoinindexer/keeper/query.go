package keeper

import (
	"github.com/evmos/ethermint/x/bitcoinindexer/types"
)

var _ types.QueryServer = Keeper{}
