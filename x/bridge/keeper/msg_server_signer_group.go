package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/evmos/ethermint/x/bridge/types"
)

func (k msgServer) CreateSignerGroup(goCtx context.Context, msg *types.MsgCreateSignerGroup) (*types.MsgCreateSignerGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetSignerGroup(
		ctx,
		msg.Name,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var signerGroup = types.SignerGroup{
		Creator: msg.Creator,
		Name:    msg.Name,
		Admin:   msg.Admin,
		Members: msg.Members,
	}

	k.SetSignerGroup(
		ctx,
		signerGroup,
	)
	return &types.MsgCreateSignerGroupResponse{}, nil
}

func (k msgServer) UpdateSignerGroup(goCtx context.Context, msg *types.MsgUpdateSignerGroup) (*types.MsgUpdateSignerGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetSignerGroup(
		ctx,
		msg.Name,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var signerGroup = types.SignerGroup{
		Creator: msg.Creator,
		Name:    msg.Name,
		Admin:   msg.Admin,
		Members: msg.Members,
	}

	k.SetSignerGroup(ctx, signerGroup)

	return &types.MsgUpdateSignerGroupResponse{}, nil
}

func (k msgServer) DeleteSignerGroup(goCtx context.Context, msg *types.MsgDeleteSignerGroup) (*types.MsgDeleteSignerGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetSignerGroup(
		ctx,
		msg.Name,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveSignerGroup(
		ctx,
		msg.Name,
	)

	return &types.MsgDeleteSignerGroupResponse{}, nil
}
