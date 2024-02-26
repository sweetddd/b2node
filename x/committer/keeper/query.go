package keeper

import (
	"github.com/evmos/ethermint/x/committer/types"
)

var _ types.QueryServer = Keeper{}
