package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/evmos/ethermint/x/bridge/types"
)

func (k msgServer) CreateWithdraw(goCtx context.Context, msg *types.MsgCreateWithdraw) (*types.MsgCreateWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the sender is in caller group.
	params := k.GetParams(ctx)
	if !k.IsMemberInCallerGroup(ctx, params.GetCallerGroupName(), msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "sender is not in caller group")
	}
	// Check if the value already exists
	_, isFound := k.GetWithdraw(
		ctx,
		msg.TxHash,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var withdraw = types.Withdraw{
		Creator:    msg.Creator,
		TxHash:     msg.TxHash,
		From:       msg.From,
		To:         msg.To,
		CoinType:   msg.CoinType,
		Value:      msg.Value,
		Data:       msg.Data,
		Status:     "pending",
		Signatures: []string{},
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
		msg.TxHash,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}
	if valFound.GetStatus() != "signed" {
		return nil, sdkerrors.Wrap(types.ErrInvalidStatus, "status is not signed")
	}

	// Check if the sender is in caller group.
	params := k.GetParams(ctx)
	if !k.IsMemberInCallerGroup(ctx, params.GetCallerGroupName(), msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "sender is not in caller group")
	}

	var withdraw = types.Withdraw{
		Creator:    valFound.Creator,
		TxHash:     valFound.TxHash,
		From:       valFound.From,
		To:         valFound.To,
		CoinType:   valFound.CoinType,
		Value:      valFound.Value,
		Data:       valFound.Data,
		Status:     msg.Status,
		Signatures: valFound.Signatures,
	}

	k.SetWithdraw(ctx, withdraw)

	return &types.MsgUpdateWithdrawResponse{}, nil
}

func (k msgServer) SignWithdraw(goCtx context.Context, msg *types.MsgSignWithdraw) (*types.MsgSignWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetWithdraw(
		ctx,
		msg.TxHash,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}
	if valFound.GetStatus() != "pending" {
		return nil, sdkerrors.Wrap(types.ErrInvalidStatus, "status is not pending")
	}

	// Check if the sender is in caller group.
	params := k.GetParams(ctx)
	if !k.IsMemberInSignerGroup(ctx, params.GetSignerGroupName(), msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "sender is not in signer group")
	}

	// if len(signatures) >= 3, Change withdraw status.
	signatures := append(valFound.GetSignatures(), msg.Signature)
	status := valFound.Status
	if len(signatures) >= 3 {
		status = "signed"
	}

	var withdraw = types.Withdraw{
		Creator:    valFound.Creator,
		TxHash:     valFound.TxHash,
		From:       valFound.From,
		To:         valFound.To,
		CoinType:   valFound.CoinType,
		Value:      valFound.Value,
		Data:       valFound.Data,
		Status:     status,
		Signatures: signatures,
	}

	k.SetWithdraw(ctx, withdraw)

	return &types.MsgSignWithdrawResponse{}, nil
}

func (k msgServer) DeleteWithdraw(goCtx context.Context, msg *types.MsgDeleteWithdraw) (*types.MsgDeleteWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetWithdraw(
		ctx,
		msg.TxHash,
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
		msg.TxHash,
	)

	return &types.MsgDeleteWithdrawResponse{}, nil
}
