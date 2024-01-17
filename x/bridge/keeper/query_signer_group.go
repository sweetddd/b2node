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

func (k Keeper) SignerGroupAll(goCtx context.Context, req *types.QueryAllSignerGroupRequest) (*types.QueryAllSignerGroupResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var signerGroups []types.SignerGroup
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	signerGroupStore := prefix.NewStore(store, types.KeyPrefix(types.SignerGroupKeyPrefix))

	pageRes, err := query.Paginate(signerGroupStore, req.Pagination, func(key []byte, value []byte) error {
		var signerGroup types.SignerGroup
		if err := k.cdc.Unmarshal(value, &signerGroup); err != nil {
			return err
		}

		signerGroups = append(signerGroups, signerGroup)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSignerGroupResponse{SignerGroup: signerGroups, Pagination: pageRes}, nil
}

func (k Keeper) SignerGroup(goCtx context.Context, req *types.QueryGetSignerGroupRequest) (*types.QueryGetSignerGroupResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetSignerGroup(
		ctx,
		req.Name,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSignerGroupResponse{SignerGroup: val}, nil
}
