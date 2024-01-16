package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/evmos/ethermint/x/bridge/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CallerAll(goCtx context.Context, req *types.QueryAllCallerRequest) (*types.QueryAllCallerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var callers []types.Caller
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	callerStore := prefix.NewStore(store, types.KeyPrefix(types.CallerKey))

	pageRes, err := query.Paginate(callerStore, req.Pagination, func(key []byte, value []byte) error {
		var caller types.Caller
		if err := k.cdc.Unmarshal(value, &caller); err != nil {
			return err
		}

		callers = append(callers, caller)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCallerResponse{Caller: callers, Pagination: pageRes}, nil
}

func (k Keeper) Caller(goCtx context.Context, req *types.QueryGetCallerRequest) (*types.QueryGetCallerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	caller, found := k.GetCaller(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetCallerResponse{Caller: caller}, nil
}
