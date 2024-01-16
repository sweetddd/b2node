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

func (k Keeper) SignerAll(goCtx context.Context, req *types.QueryAllSignerRequest) (*types.QueryAllSignerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var signers []types.Signer
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	signerStore := prefix.NewStore(store, types.KeyPrefix(types.SignerKey))

	pageRes, err := query.Paginate(signerStore, req.Pagination, func(key []byte, value []byte) error {
		var signer types.Signer
		if err := k.cdc.Unmarshal(value, &signer); err != nil {
			return err
		}

		signers = append(signers, signer)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSignerResponse{Signer: signers, Pagination: pageRes}, nil
}

func (k Keeper) Signer(goCtx context.Context, req *types.QueryGetSignerRequest) (*types.QueryGetSignerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	signer, found := k.GetSigner(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSignerResponse{Signer: signer}, nil
}
