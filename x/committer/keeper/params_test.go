package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/evmos/ethermint/testutil/keeper"
	"github.com/evmos/ethermint/x/committer/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CommitterKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
