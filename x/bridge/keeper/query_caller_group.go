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

func (k Keeper) CallerGroupAll(goCtx context.Context, req *types.QueryAllCallerGroupRequest) (*types.QueryAllCallerGroupResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var callerGroups []types.CallerGroup
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	callerGroupStore := prefix.NewStore(store, types.KeyPrefix(types.CallerGroupKeyPrefix))

	pageRes, err := query.Paginate(callerGroupStore, req.Pagination, func(key []byte, value []byte) error {
		var callerGroup types.CallerGroup
		if err := k.cdc.Unmarshal(value, &callerGroup); err != nil {
			return err
		}

		callerGroups = append(callerGroups, callerGroup)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCallerGroupResponse{CallerGroup: callerGroups, Pagination: pageRes}, nil
}

func (k Keeper) CallerGroup(goCtx context.Context, req *types.QueryGetCallerGroupRequest) (*types.QueryGetCallerGroupResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetCallerGroup(
		ctx,
		req.Name,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCallerGroupResponse{CallerGroup: val}, nil
}
