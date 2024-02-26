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
		msg.TxId,
	)
	if isFound {
		return nil, types.ErrIndexExist
	}
	for _, txHash := range msg.TxHashList {
		_, isTxHashFound := k.GetRollupTx(ctx, txHash)
		if isTxHashFound {
			return nil, types.ErrRollupTxHashExist
		}
	}
	for _, txHash := range msg.TxHashList {
		k.SetRollupTx(ctx, types.RollupTx{
			TxHash: txHash,
			TxId:   msg.TxId,
			Status: types.WithdrawStatus_WITHDRAW_STATUS_PENDING,
		})
	}

	withdraw := types.Withdraw{
		Creator:     msg.Creator,
		TxId:        msg.TxId,
		TxHashList:  msg.TxHashList,
		EncodedData: msg.EncodedData,
		Status:      types.WithdrawStatus_WITHDRAW_STATUS_PENDING,
		Signatures:  make(map[string]string),
	}

	k.SetWithdraw(
		ctx,
		withdraw,
	)

	k.SetStatusIndex(ctx, withdraw.Status.String(), withdraw.TxId)

	if err := ctx.EventManager().EmitTypedEvent(&types.EventCreateWithdraw{TxId: msg.TxId}); err != nil {
		return nil, err
	}

	return &types.MsgCreateWithdrawResponse{}, nil
}

func (k msgServer) UpdateWithdraw(goCtx context.Context, msg *types.MsgUpdateWithdraw) (*types.MsgUpdateWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetWithdraw(
		ctx,
		msg.TxId,
	)
	if !isFound {
		return nil, types.ErrIndexNotExist
	}
	if valFound.GetStatus() != types.WithdrawStatus_WITHDRAW_STATUS_SIGNED {
		return nil, types.ErrInvalidStatus
	}

	// Check if the sender is in caller group.
	params := k.GetParams(ctx)
	if !k.IsMemberInCallerGroup(ctx, params.GetCallerGroupName(), msg.Creator) {
		return nil, types.ErrNotCallerGroupMembers
	}

	for _, txHash := range valFound.TxHashList {
		k.SetRollupTx(ctx, types.RollupTx{
			TxHash: txHash,
			TxId:   valFound.TxId,
			Status: msg.Status,
		})
	}

	withdraw := types.Withdraw{
		Creator:     valFound.Creator,
		TxId:        valFound.TxId,
		TxHashList:  valFound.TxHashList,
		EncodedData: valFound.EncodedData,
		Status:      msg.Status,
		Signatures:  valFound.Signatures,
	}

	k.SetWithdraw(ctx, withdraw)
	k.RemoveStatusIndex(ctx, valFound.GetStatus().String(), valFound.TxId)
	k.SetStatusIndex(ctx, withdraw.Status.String(), withdraw.TxId)

	if err := ctx.EventManager().EmitTypedEvent(&types.EventUpdateWithdraw{TxId: msg.TxId}); err != nil {
		return nil, err
	}

	return &types.MsgUpdateWithdrawResponse{}, nil
}

func (k msgServer) SignWithdraw(goCtx context.Context, msg *types.MsgSignWithdraw) (*types.MsgSignWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetWithdraw(
		ctx,
		msg.TxId,
	)
	if !isFound {
		return nil, types.ErrIndexNotExist
	}
	if valFound.GetStatus() != types.WithdrawStatus_WITHDRAW_STATUS_PENDING {
		return nil, types.ErrInvalidStatus
	}

	// Check if the sender is in caller group.
	params := k.GetParams(ctx)
	signerGroupName := params.GetSignerGroupName()
	if !k.IsMemberInSignerGroup(ctx, signerGroupName, msg.Creator) {
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

	threshold := int(k.GetSignerGroupThreshold(ctx, signerGroupName)) // #nosec
	if threshold == 0 {
		return nil, types.ErrThresholdNotSet
	}
	if len(signatures) >= threshold {
		status = types.WithdrawStatus_WITHDRAW_STATUS_SIGNED
	}

	withdraw := types.Withdraw{
		Creator:     valFound.Creator,
		TxId:        valFound.TxId,
		TxHashList:  valFound.TxHashList,
		EncodedData: valFound.EncodedData,
		Status:      status,
		Signatures:  signatures,
	}

	k.SetWithdraw(ctx, withdraw)

	if status == types.WithdrawStatus_WITHDRAW_STATUS_SIGNED {
		if err := ctx.EventManager().EmitTypedEvent(&types.EventSignWithdraw{TxId: msg.TxId}); err != nil {
			return nil, err
		}
		for _, txHash := range valFound.TxHashList {
			k.SetRollupTx(ctx, types.RollupTx{
				TxHash: txHash,
				TxId:   valFound.TxId,
				Status: status,
			})
		}
		k.RemoveStatusIndex(ctx, valFound.GetStatus().String(), valFound.TxId)
		k.SetStatusIndex(ctx, withdraw.Status.String(), withdraw.TxId)
	}
	return &types.MsgSignWithdrawResponse{}, nil
}

func (k msgServer) DeleteWithdraw(goCtx context.Context, msg *types.MsgDeleteWithdraw) (*types.MsgDeleteWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetWithdraw(
		ctx,
		msg.TxId,
	)
	if !isFound {
		return nil, types.ErrIndexNotExist
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, types.ErrNotOwner
	}

	for _, txHash := range valFound.TxHashList {
		k.RemoveRollupTx(ctx, txHash)
	}

	k.RemoveWithdraw(
		ctx,
		msg.TxId,
	)

	k.RemoveStatusIndex(ctx, valFound.GetStatus().String(), valFound.TxId)

	if err := ctx.EventManager().EmitTypedEvent(&types.EventDeleteWithdraw{TxId: msg.TxId}); err != nil {
		return nil, err
	}

	return &types.MsgDeleteWithdrawResponse{}, nil
}
