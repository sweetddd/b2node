package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/x/committer/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// SubmitProof defines the rpc handler for MsgSubmitProof.
func (k msgServer) SubmitProof(goCtx context.Context, msg *types.MsgSubmitProof) (*types.MsgSubmitProofResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check committer permission
	if !k.IsExistCommitter(ctx, msg.From) {
		return &types.MsgSubmitProofResponse{}, types.ErrAccountPermission
	}

	proposal, found := k.GetProposal(ctx, msg.Id)
	// If proposal not found, create a new one
	if !found {
		lastProposal := k.GetLastProposal(ctx)
		proposal = types.Proposal{
			Id:            lastProposal.Id + 1,
			Proposer:      msg.From,
			ProofHash:     msg.ProofHash,
			StateRootHash: msg.StateRootHash,
			StartIndex:    msg.StartIndex,
			EndIndex:      msg.EndIndex,
			BlockHight:    uint64(ctx.BlockHeight()),
			Status:        types.VotingStatus,
		}

		if !(lastProposal.EndIndex == 0 && proposal.StartIndex == 1) &&
			!(lastProposal.EndIndex != 0 && lastProposal.EndIndex == proposal.StartIndex) {
			return &types.MsgSubmitProofResponse{},
				fmt.Errorf(
					"proposal start index must equal last proposal end index, "+
						"last proposal end index: %s", fmt.Sprint(lastProposal.EndIndex))
		}

		k.SetProposal(ctx, proposal)
		k.SetLastProposal(ctx, proposal)
	}

	if proposal.Status != types.VotingStatus {
		return &types.MsgSubmitProofResponse{}, types.ErrProposalStatus
	}

	if k.CheckAndUpdateProposalTimeout(ctx, proposal) {
		return &types.MsgSubmitProofResponse{}, types.ErrProposalTimeout
	}

	if k.HasVoted(msg.From, proposal.VotedListPhaseCommit) {
		return &types.MsgSubmitProofResponse{}, types.ErrAlreadyVoted
	}

	// Vote for the proposal and update status
	k.VoteAndUpdateProposal(ctx, proposal, msg.From)

	return &types.MsgSubmitProofResponse{Id: msg.Id}, nil
}

// BitcoinTx defines the rpc handler for MsgBitcoinTx.
func (k msgServer) BitcoinTx(goCtx context.Context, msg *types.MsgBitcoinTx) (*types.MsgBitcoinTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check committer permission
	if !k.IsExistCommitter(ctx, msg.From) {
		return &types.MsgBitcoinTxResponse{}, types.ErrAccountPermission
	}

	proposal, found := k.GetProposal(ctx, msg.Id)
	if !found {
		return &types.MsgBitcoinTxResponse{}, types.ErrNotExistProposal
	}

	if proposal.Status != types.PendingStatus {
		return &types.MsgBitcoinTxResponse{}, types.ErrProposalStatus
	}

	if proposal.BitcoinTxHash == "" && proposal.Winner != msg.From {
		return &types.MsgBitcoinTxResponse{}, types.ErrAccountPermission
	}

	if k.HasVoted(msg.From, proposal.VotedListPhaseTimeout) {
		return &types.MsgBitcoinTxResponse{}, types.ErrAlreadyVoted
	}

	if k.CheckAndUpdateProposalTimeout(ctx, proposal) {
		return &types.MsgBitcoinTxResponse{}, types.ErrProposalTimeout
	}

	k.VoteAndUpdateBitcoinTx(ctx, proposal, msg.From, msg.BitcoinTxHash)

	return &types.MsgBitcoinTxResponse{Id: proposal.Id}, nil
}

// TimeoutProposal defines the rpc handler for MsgTimeoutProposal.
func (k msgServer) TimeoutProposal(goCtx context.Context, msg *types.MsgTimeoutProposal) (*types.MsgTimeoutProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check committer permission
	if !k.IsExistCommitter(ctx, msg.From) {
		return &types.MsgTimeoutProposalResponse{}, types.ErrAccountPermission
	}

	proposal, found := k.GetProposal(ctx, msg.Id)
	if !found {
		return &types.MsgTimeoutProposalResponse{}, types.ErrNotExistProposal
	}

	if proposal.Status != types.PendingStatus {
		return &types.MsgTimeoutProposalResponse{}, types.ErrProposalStatus
	}

	isTimeout := k.CheckAndUpdateProposalTimeout(ctx, proposal)
	if !isTimeout {
		return &types.MsgTimeoutProposalResponse{}, types.ErrInvalidProposal
	}

	return &types.MsgTimeoutProposalResponse{}, nil
}

// AddCommitter defines the rpc handler for MsgAddCommitter.
func (k msgServer) AddCommitter(goCtx context.Context, msg *types.MsgAddCommitter) (*types.MsgAddCommitterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check admin permission
	if msg.From != k.GetParams(ctx).GetAdminPolicyAccount(types.PolicyType_POLICY_TYPE_GROUP1) {
		return &types.MsgAddCommitterResponse{}, types.ErrAccountPermission
	}

	found := k.IsExistCommitter(ctx, msg.Committer)
	if found {
		return &types.MsgAddCommitterResponse{}, types.ErrExistCommitter
	}

	committers := k.GetAllCommitters(ctx)
	committers.CommitterList = append(committers.CommitterList, msg.Committer)
	k.SetCommitter(ctx, committers)

	return &types.MsgAddCommitterResponse{Committer: msg.Committer}, nil
}

// RemoveCommitter defines the rpc handler for MsgRemoveCommitter.
func (k msgServer) RemoveCommitter(goCtx context.Context, msg *types.MsgRemoveCommitter) (*types.MsgRemoveCommitterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check admin permission
	if msg.From != k.GetParams(ctx).GetAdminPolicyAccount(types.PolicyType_POLICY_TYPE_GROUP1) {
		return &types.MsgRemoveCommitterResponse{}, types.ErrAccountPermission
	}

	found := k.IsExistCommitter(ctx, msg.Committer)
	if !found {
		return &types.MsgRemoveCommitterResponse{}, types.ErrNotExistCommitter
	}

	committers := k.GetAllCommitters(ctx)
	for i, committer := range committers.CommitterList {
		if committer == msg.Committer {
			committers.CommitterList = append(committers.CommitterList[:i], committers.CommitterList[i+1:]...)
			break
		}
	}
	k.SetCommitter(ctx, committers)

	return &types.MsgRemoveCommitterResponse{Committer: msg.Committer}, nil
}
