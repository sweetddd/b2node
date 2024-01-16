package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/evmos/ethermint/x/bridge/types"
)

func (k msgServer) CreateSigner(goCtx context.Context, msg *types.MsgCreateSigner) (*types.MsgCreateSignerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var signer = types.Signer{
		Creator: msg.Creator,
		Address: msg.Address,
		Name:    msg.Name,
	}

	id := k.AppendSigner(
		ctx,
		signer,
	)

	return &types.MsgCreateSignerResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateSigner(goCtx context.Context, msg *types.MsgUpdateSigner) (*types.MsgUpdateSignerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var signer = types.Signer{
		Creator: msg.Creator,
		Id:      msg.Id,
		Address: msg.Address,
		Name:    msg.Name,
	}

	// Checks that the element exists
	val, found := k.GetSigner(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetSigner(ctx, signer)

	return &types.MsgUpdateSignerResponse{}, nil
}

func (k msgServer) DeleteSigner(goCtx context.Context, msg *types.MsgDeleteSigner) (*types.MsgDeleteSignerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetSigner(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveSigner(ctx, msg.Id)

	return &types.MsgDeleteSignerResponse{}, nil
}
