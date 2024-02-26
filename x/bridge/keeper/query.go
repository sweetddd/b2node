package keeper

import (
	"github.com/evmos/ethermint/x/bridge/types"
)

var _ types.QueryServer = Keeper{}
