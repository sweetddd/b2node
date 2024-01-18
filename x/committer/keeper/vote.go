package keeper

import (
	types "github.com/evmos/ethermint/x/committer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) VoteAndUpdateProposal(ctx sdk.Context, proposal types.Proposal, from string) {
	proposal.VotedListPhaseCommit = append(proposal.VotedListPhaseCommit, from)
	allCommitter := k.GetAllCommitters(ctx)
	if len(allCommitter) / 2 < len(proposal.VotedListPhaseCommit) {
		proposal.Status = types.Pending_Status
		winnerIndex := ctx.BlockHeight() % int64(len(proposal.VotedListPhaseCommit))
		proposal.Winner = proposal.VotedListPhaseCommit[winnerIndex]
	}
	k.SetProposal(ctx, proposal)
}

func (k Keeper) CheckAndUpdateProposalTimeout(ctx sdk.Context, proposal types.Proposal) bool {
	if timeout := k.IsTimeout(ctx, proposal); timeout {
		proposal.Status = types.Timeout_Status
		k.SetProposal(ctx, proposal)
		return true
	}

	return false
}

func (k Keeper) IsTimeout(ctx sdk.Context, proposal types.Proposal) bool {
	currBlockHight := ctx.BlockHeight();
	return currBlockHight - int64(proposal.BlockHight) > types.DefaultTimeoutBlockPeriod
}