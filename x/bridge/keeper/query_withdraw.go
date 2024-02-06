package keeper

import (
	"context"
	"fmt"

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
		req.TxId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetWithdrawResponse{Withdraw: val}, nil
}

func (k Keeper) WithdrawsByStatus(goCtx context.Context, req *types.QueryWithdrawsByStatusRequest) (*types.QueryWithdrawsByStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	withdraws := k.GetAllWithdrawByStatus(ctx, req.Status.String())
	paginatedWithdraws, pageRes, err := PaginateWithdraws(withdraws, req.Pagination)
	if err != nil {
		return nil, err
	}

	return &types.QueryWithdrawsByStatusResponse{
		Withdraw:   paginatedWithdraws,
		Pagination: pageRes,
	}, nil
}

// PaginateWithdraws is a helper function to do pagination of withdraws.
func PaginateWithdraws(withdraws []types.Withdraw, pageReq *query.PageRequest) ([]types.Withdraw, *query.PageResponse, error) {
	if pageReq == nil {
		pageReq = &query.PageRequest{
			Limit:  100, // 设定一个默认值或者根据具体情况处理
			Offset: 0,
		}
	}

	start, end := int(pageReq.Offset), int(pageReq.Offset+pageReq.Limit) // #nosec
	if end > len(withdraws) || end < 0 {                                 // -ve end means that the end has overflown
		end = len(withdraws)
	}
	if start > len(withdraws) || start < 0 {
		return []types.Withdraw{}, nil, status.Errorf(codes.InvalidArgument, "invalid request")
	}

	paginatedWithdraws := withdraws[start:end]

	var nextPageToken []byte
	if end < len(withdraws) {
		nextPageToken = []byte(fmt.Sprintf("%v", end)) // Serialize the end cursor
	}

	return paginatedWithdraws, &query.PageResponse{
		NextKey: nextPageToken,
		Total:   uint64(len(withdraws)),
	}, nil
}
