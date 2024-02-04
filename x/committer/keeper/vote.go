package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	types "github.com/evmos/ethermint/x/committer/types"
)

func (k Keeper) VoteAndUpdateProposal(ctx sdk.Context, proposal types.Proposal, from string) {
	proposal.VotedListPhaseCommit = append(proposal.VotedListPhaseCommit, from)
	allCommitter := k.GetAllCommitters(ctx)
	if len(allCommitter.CommitterList)/2 < len(proposal.VotedListPhaseCommit) {
		proposal.Status = types.PendingStatus
		winnerIndex := ctx.BlockHeight() % int64(len(proposal.VotedListPhaseCommit))
		proposal.Winner = proposal.VotedListPhaseCommit[winnerIndex]
	}
	k.SetProposal(ctx, proposal)
}

func (k Keeper) VoteAndUpdateBitcoinTx(ctx sdk.Context, proposal types.Proposal, from string, bitcoinTx string) {
	if proposal.BitcoinTxHash == "" {
		proposal.BitcoinTxHash = bitcoinTx
	}
	proposal.VotedListPhaseTimeout = append(proposal.VotedListPhaseTimeout, from)
	allCommitter := k.GetAllCommitters(ctx)
	if len(allCommitter.CommitterList)/2 < len(proposal.VotedListPhaseTimeout) {
		proposal.Status = types.SucceedStatus
	}
	k.SetProposal(ctx, proposal)
}

func (k Keeper) CheckAndUpdateProposalTimeout(ctx sdk.Context, proposal types.Proposal) bool {
	if timeout := k.IsTimeout(ctx, proposal); timeout {
		proposal.Status = types.TimeoutStatus
		k.SetProposal(ctx, proposal)
		return true
	}

	return false
}

func (k Keeper) HasVoted(address string, votedList []string) bool {
	for _, voted := range votedList {
		if voted == address {
			return true
		}
	}

	return false
}

func (k Keeper) IsTimeout(ctx sdk.Context, proposal types.Proposal) bool {
	currBlockHight := uint64(ctx.BlockHeight())
	return currBlockHight-proposal.BlockHight > types.DefaultTimeoutBlockPeriod
}
