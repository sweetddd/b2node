package keeper

import (
	"context"
	"github.com/evmos/ethermint/x/committer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BatchProof defines the rpc handler for MsgBatchProof.
func (k msgServer) BatchProof(goCtx context.Context, msg *types.MsgBatchProofTx) (*types.MsgBatchProofTxResponse, error) {
	// TODO: check permission
	ctx := sdk.UnwrapSDKContext(goCtx)

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
	// TODO: check permission
	ctx := sdk.UnwrapSDKContext(goCtx)

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
	// TODO: check permission
	ctx := sdk.UnwrapSDKContext(goCtx)

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

