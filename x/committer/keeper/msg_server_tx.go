package keeper

import (
	"context"
	"github.com/evmos/ethermint/x/committer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BatchProof defines the rpc handler for MsgBatchProof.
func (k msgServer) BatchProof(goCtx context.Context, msg *types.MsgBatchProofTx) (*types.MsgBatchProofTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check committer permission
	if !k.IsExistCommitter(ctx, msg.From) {
		return &types.MsgBatchProofTxResponse{}, types.ErrAccountPermission
	}

	proposal, found := k.GetProposal(ctx, msg.Id)
	// If proposal not found, create a new one
	if !found {
		proposalId := k.GetLastProposal(ctx).Id + 1;
		proposal = types.Proposal{
			Id: proposalId,
			Proposer: msg.From,
			ProofHash: msg.ProofHash,
			StateRootHash: msg.StateRootHash,
			StartIndex: msg.StartIndex,
			EndIndex: msg.EndIndex,
			BlockHight: uint64(ctx.BlockHeight()),
			Status: types.Voting_Status,
		}

		k.SetProposal(ctx, proposal)
	}

	if proposal.Status != types.Voting_Status {
		return &types.MsgBatchProofTxResponse{}, types.ErrProposalStatus
	}

	if k.CheckAndUpdateProposalTimeout(ctx, proposal) {
		return &types.MsgBatchProofTxResponse{}, types.ErrProposalTimeout
	}

	// Vote for the proposal and update status
	k.VoteAndUpdateProposal(ctx, proposal, msg.From)

	return &types.MsgBatchProofTxResponse{Id: msg.Id}, nil
}

// TapRoot defines the rpc handler for MsgTapRoot.	
func (k msgServer) TapRoot(goCtx context.Context, msg *types.MsgTapRootTx) (*types.MsgTapRootTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check committer permission
	if !k.IsExistCommitter(ctx, msg.From) {
		return &types.MsgTapRootTxResponse{}, types.ErrAccountPermission
	}

	proposal, found := k.GetProposal(ctx, msg.Id)
	if !found {
		return &types.MsgTapRootTxResponse{}, types.ErrNotExistProposal
	}

	if proposal.Winner != msg.From {
		return &types.MsgTapRootTxResponse{}, types.ErrAccountPermission
	}

	if proposal.Status != types.Pending_Status {
		return &types.MsgTapRootTxResponse{}, types.ErrProposalStatus
	}

	if k.CheckAndUpdateProposalTimeout(ctx, proposal) {
		return &types.MsgTapRootTxResponse{}, types.ErrProposalTimeout
	}

	proposal.BitcoinTxHash = msg.BitcoinTxHash
	k.SetProposal(ctx, proposal)

	return &types.MsgTapRootTxResponse{Id: proposal.Id}, nil
}

// TimeoutProposal defines the rpc handler for MsgTimeoutProposal.
func (k msgServer) TimeoutProposal(goCtx context.Context, msg *types.MsgTimeoutProposalTx) (*types.MsgTimeoutProposalTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check committer permission
	if !k.IsExistCommitter(ctx, msg.From) {
		return &types.MsgTimeoutProposalTxResponse{}, types.ErrAccountPermission
	}

	proposal, found := k.GetProposal(ctx, msg.Id)
	if !found {
		return &types.MsgTimeoutProposalTxResponse{}, types.ErrNotExistProposal
	}

	if proposal.Status != types.Pending_Status {
		return &types.MsgTimeoutProposalTxResponse{}, types.ErrProposalStatus
	}

	isTimeout := k.CheckAndUpdateProposalTimeout(ctx, proposal)
	if !isTimeout {
		return &types.MsgTimeoutProposalTxResponse{}, types.ErrInvalidProposal
	}
	
	return &types.MsgTimeoutProposalTxResponse{}, nil
}

// AddCommitter defines the rpc handler for MsgAddCommitter.
func (k msgServer) AddCommitter(goCtx context.Context, msg *types.MsgAddCommitterTx) (*types.MsgAddCommitterTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check admin permission
	if msg.From != k.GetParams(ctx).GetAdminPolicyAccount(types.PolicyType_group1) {
		return &types.MsgAddCommitterTxResponse{}, types.ErrAccountPermission
	}
	
	found := k.IsExistCommitter(ctx, msg.Committer)
	if found {
		return &types.MsgAddCommitterTxResponse{}, types.ErrExistCommitter
	}

	committers := k.GetAllCommitters(ctx)
	committers.CommitterList = append(committers.CommitterList, msg.Committer)
	k.SetCommitter(ctx, committers)

	return &types.MsgAddCommitterTxResponse{Committer: msg.Committer}, nil
}

// RemoveCommitter defines the rpc handler for MsgRemoveCommitter.
func (k msgServer) RemoveCommitter(goCtx context.Context, msg *types.MsgRemoveCommitterTx) (*types.MsgRemoveCommitterTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check admin permission
	if msg.From != k.GetParams(ctx).GetAdminPolicyAccount(types.PolicyType_group1) {
		return &types.MsgRemoveCommitterTxResponse{}, types.ErrAccountPermission
	}

	found := k.IsExistCommitter(ctx, msg.Committer)
	if !found {
		return &types.MsgRemoveCommitterTxResponse{}, types.ErrNotExistCommitter
	}

	committers := k.GetAllCommitters(ctx)
	for i, committer := range committers.CommitterList {
		if committer == msg.Committer {
			committers.CommitterList = append(committers.CommitterList[:i], committers.CommitterList[i+1:]...)
			break
		}
	}
	k.SetCommitter(ctx, committers)

	return &types.MsgRemoveCommitterTxResponse{Committer: msg.Committer}, nil
}