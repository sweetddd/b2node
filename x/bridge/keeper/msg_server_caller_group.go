package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/evmos/ethermint/x/bridge/types"
)

func (k msgServer) CreateCallerGroup(goCtx context.Context, msg *types.MsgCreateCallerGroup) (*types.MsgCreateCallerGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetCallerGroup(
		ctx,
		msg.Name,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var callerGroup = types.CallerGroup{
		Creator: msg.Creator,
		Name:    msg.Name,
		Admin:   msg.Admin,
		Members: msg.Members,
	}

	k.SetCallerGroup(
		ctx,
		callerGroup,
	)
	return &types.MsgCreateCallerGroupResponse{}, nil
}

func (k msgServer) UpdateCallerGroup(goCtx context.Context, msg *types.MsgUpdateCallerGroup) (*types.MsgUpdateCallerGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetCallerGroup(
		ctx,
		msg.Name,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current admin
	if msg.Creator != valFound.GetAdmin() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var callerGroup = types.CallerGroup{
		Creator: valFound.Creator,
		Name:    msg.Name,
		Admin:   msg.Admin,
		Members: msg.Members,
	}

	k.SetCallerGroup(ctx, callerGroup)

	return &types.MsgUpdateCallerGroupResponse{}, nil
}

func (k msgServer) DeleteCallerGroup(goCtx context.Context, msg *types.MsgDeleteCallerGroup) (*types.MsgDeleteCallerGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetCallerGroup(
		ctx,
		msg.Name,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.GetAdmin() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveCallerGroup(
		ctx,
		msg.Name,
	)

	return &types.MsgDeleteCallerGroupResponse{}, nil
}
