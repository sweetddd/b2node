package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/evmos/ethermint/testutil/bridge/keeper"
	"github.com/evmos/ethermint/testutil/bridge/nullify"
	"github.com/evmos/ethermint/x/bridge/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestRollupTxQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRollupTx(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetRollupTxRequest
		response *types.QueryGetRollupTxResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetRollupTxRequest{
				TxHash: msgs[0].TxHash,
			},
			response: &types.QueryGetRollupTxResponse{RollupTx: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetRollupTxRequest{
				TxHash: msgs[1].TxHash,
			},
			response: &types.QueryGetRollupTxResponse{RollupTx: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetRollupTxRequest{
				TxHash: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.RollupTx(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestRollupTxQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.BridgeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRollupTx(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRollupTxRequest {
		return &types.QueryAllRollupTxRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RollupTxAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RollupTx), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RollupTx),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RollupTxAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RollupTx), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RollupTx),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.RollupTxAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.RollupTx),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.RollupTxAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
