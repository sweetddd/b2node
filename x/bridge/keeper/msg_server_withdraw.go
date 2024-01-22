package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/x/bridge/types"
)

func (k msgServer) CreateWithdraw(goCtx context.Context, msg *types.MsgCreateWithdraw) (*types.MsgCreateWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the sender is in caller group.
	params := k.GetParams(ctx)
	if !k.IsMemberInCallerGroup(ctx, params.GetCallerGroupName(), msg.Creator) {
		return nil, types.ErrNotCallerGroupMembers
	}
	// Check if the value already exists
	_, isFound := k.GetWithdraw(
		ctx,
		msg.TxHash,
	)
	if isFound {
		return nil, types.ErrIndexExist
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
		Signatures: make(map[string]string),
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
		return nil, types.ErrIndexNotExist
	}
	if valFound.GetStatus() != "signed" {
		return nil, types.ErrInvalidStatus
	}

	// Check if the sender is in caller group.
	params := k.GetParams(ctx)
	if !k.IsMemberInCallerGroup(ctx, params.GetCallerGroupName(), msg.Creator) {
		return nil, types.ErrNotCallerGroupMembers
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
		return nil, types.ErrIndexNotExist
	}
	if valFound.GetStatus() != "pending" {
		return nil, types.ErrInvalidStatus
	}

	// Check if the sender is in caller group.
	params := k.GetParams(ctx)
	if !k.IsMemberInSignerGroup(ctx, params.GetSignerGroupName(), msg.Creator) {
		return nil, types.ErrNotSignerGroupMembers
	}

	signatures := valFound.GetSignatures()
	if signatures == nil {
		signatures = make(map[string]string)
	} else {
		_, ok := signatures[msg.Creator]
		if ok {
			return nil, types.ErrAlreadySigned
		}
	}
	signatures[msg.Creator] = msg.Signature
	// if len(signatures) >= 3, Change withdraw status.
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
		return nil, types.ErrIndexNotExist
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, types.ErrNotOwner
	}

	k.RemoveWithdraw(
		ctx,
		msg.TxHash,
	)

	return &types.MsgDeleteWithdrawResponse{}, nil
}
