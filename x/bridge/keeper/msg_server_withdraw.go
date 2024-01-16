package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/evmos/ethermint/x/bridge/types"
)

func (k msgServer) CreateWithdraw(goCtx context.Context, msg *types.MsgCreateWithdraw) (*types.MsgCreateWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetWithdraw(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var withdraw = types.Withdraw{
		Creator:    msg.Creator,
		Index:      msg.Index,
		TxHash:     msg.TxHash,
		From:       msg.From,
		To:         msg.To,
		CoinType:   msg.CoinType,
		Value:      msg.Value,
		Data:       msg.Data,
		Status:     msg.Status,
		Signatures: msg.Signatures,
	}

	k.SetWithdraw(
		ctx,
		withdraw,
	)
	return &types.MsgCreateWithdrawResponse{}, nil
}

func (k msgServer) UpdateWithdraw(goCtx context.Context, msg *types.MsgUpdateWithdraw) (*types.MsgUpdateWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetWithdraw(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var withdraw = types.Withdraw{
		Creator:    msg.Creator,
		Index:      msg.Index,
		TxHash:     msg.TxHash,
		From:       msg.From,
		To:         msg.To,
		CoinType:   msg.CoinType,
		Value:      msg.Value,
		Data:       msg.Data,
		Status:     msg.Status,
		Signatures: msg.Signatures,
	}

	k.SetWithdraw(ctx, withdraw)

	return &types.MsgUpdateWithdrawResponse{}, nil
}

func (k msgServer) DeleteWithdraw(goCtx context.Context, msg *types.MsgDeleteWithdraw) (*types.MsgDeleteWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetWithdraw(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveWithdraw(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteWithdrawResponse{}, nil
}
