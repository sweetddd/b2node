package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/evmos/ethermint/x/bridge/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WithdrawAll(goCtx context.Context, req *types.QueryAllWithdrawRequest) (*types.QueryAllWithdrawResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var withdraws []types.Withdraw
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	withdrawStore := prefix.NewStore(store, types.KeyPrefix(types.WithdrawKeyPrefix))

	pageRes, err := query.Paginate(withdrawStore, req.Pagination, func(key []byte, value []byte) error {
		var withdraw types.Withdraw
		if err := k.cdc.Unmarshal(value, &withdraw); err != nil {
			return err
		}

		withdraws = append(withdraws, withdraw)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllWithdrawResponse{Withdraw: withdraws, Pagination: pageRes}, nil
}

func (k Keeper) Withdraw(goCtx context.Context, req *types.QueryGetWithdrawRequest) (*types.QueryGetWithdrawResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetWithdraw(
		ctx,
		req.TxHash,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetWithdrawResponse{Withdraw: val}, nil
}
