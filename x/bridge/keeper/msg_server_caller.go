package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/evmos/ethermint/x/bridge/types"
)

func (k msgServer) CreateCaller(goCtx context.Context, msg *types.MsgCreateCaller) (*types.MsgCreateCallerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var caller = types.Caller{
		Creator: msg.Creator,
		Address: msg.Address,
	}

	id := k.AppendCaller(
		ctx,
		caller,
	)

	return &types.MsgCreateCallerResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateCaller(goCtx context.Context, msg *types.MsgUpdateCaller) (*types.MsgUpdateCallerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var caller = types.Caller{
		Creator: msg.Creator,
		Id:      msg.Id,
		Address: msg.Address,
	}

	// Checks that the element exists
	val, found := k.GetCaller(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetCaller(ctx, caller)

	return &types.MsgUpdateCallerResponse{}, nil
}

func (k msgServer) DeleteCaller(goCtx context.Context, msg *types.MsgDeleteCaller) (*types.MsgDeleteCallerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetCaller(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveCaller(ctx, msg.Id)

	return &types.MsgDeleteCallerResponse{}, nil
}
