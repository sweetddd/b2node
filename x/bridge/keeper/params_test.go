package keeper_test

import (
	"testing"

	keepertest "github.com/evmos/ethermint/testutil/bridge/keeper"
	"github.com/evmos/ethermint/x/bridge/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.BridgeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.CallerGroupName, k.CallerGroupName(ctx))
	require.EqualValues(t, params.SignerGroupName, k.SignerGroupName(ctx))
}
