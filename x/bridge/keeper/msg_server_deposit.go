package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/x/bridge/types"
)

func (k msgServer) CreateDeposit(goCtx context.Context, msg *types.MsgCreateDeposit) (*types.MsgCreateDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the sender is in caller group.
	params := k.GetParams(ctx)
	if !k.IsMemberInCallerGroup(ctx, params.GetCallerGroupName(), msg.Creator) {
		return nil, types.ErrNotCallerGroupMembers
	}
	// Check if the value already exists
	_, isFound := k.GetDeposit(
		ctx,
		msg.TxHash,
	)
	if isFound {
		return nil, types.ErrIndexExist
	}

	deposit := types.Deposit{
		Creator:  msg.Creator,
		TxHash:   msg.TxHash,
		From:     msg.From,
		To:       msg.To,
		CoinType: msg.CoinType,
		Value:    msg.Value,
		Data:     msg.Data,
		Status:   "pending",
	}

	k.SetDeposit(
		ctx,
		deposit,
	)

	if err := ctx.EventManager().EmitTypedEvent(&types.EventCreateDeposit{TxHash: msg.TxHash}); err != nil {
		return nil, err
	}

	return &types.MsgCreateDepositResponse{}, nil
}

func (k msgServer) UpdateDeposit(goCtx context.Context, msg *types.MsgUpdateDeposit) (*types.MsgUpdateDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDeposit(
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
	if !k.IsMemberInCallerGroup(ctx, params.GetCallerGroupName(), msg.Creator) {
		return nil, types.ErrNotCallerGroupMembers
	}

	deposit := types.Deposit{
		Creator:  valFound.Creator,
		TxHash:   valFound.TxHash,
		From:     valFound.From,
		To:       valFound.To,
		CoinType: valFound.CoinType,
		Value:    valFound.Value,
		Data:     valFound.Data,
		Status:   msg.Status,
	}

	k.SetDeposit(ctx, deposit)

	if err := ctx.EventManager().EmitTypedEvent(&types.EventUpdateDeposit{TxHash: msg.TxHash}); err != nil {
		return nil, err
	}

	return &types.MsgUpdateDepositResponse{}, nil
}

func (k msgServer) DeleteDeposit(goCtx context.Context, msg *types.MsgDeleteDeposit) (*types.MsgDeleteDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDeposit(
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

	k.RemoveDeposit(
		ctx,
		msg.TxHash,
	)

	if err := ctx.EventManager().EmitTypedEvent(&types.EventDeleteDeposit{TxHash: msg.TxHash}); err != nil {
		return nil, err
	}

	return &types.MsgDeleteDepositResponse{}, nil
}
