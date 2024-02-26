package keeper //nolint:dupl

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/evmos/ethermint/x/bridge/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RollupTxAll(goCtx context.Context, req *types.QueryAllRollupTxRequest) (*types.QueryAllRollupTxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var rollupTxs []types.RollupTx
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	rollupTxStore := prefix.NewStore(store, types.KeyPrefix(types.RollupTxKeyPrefix))

	pageRes, err := query.Paginate(rollupTxStore, req.Pagination, func(key []byte, value []byte) error {
		var rollupTx types.RollupTx
		if err := k.cdc.Unmarshal(value, &rollupTx); err != nil {
			return err
		}

		rollupTxs = append(rollupTxs, rollupTx)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRollupTxResponse{RollupTx: rollupTxs, Pagination: pageRes}, nil
}

func (k Keeper) RollupTx(goCtx context.Context, req *types.QueryGetRollupTxRequest) (*types.QueryGetRollupTxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetRollupTx(
		ctx,
		req.TxHash,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRollupTxResponse{RollupTx: val}, nil
}
