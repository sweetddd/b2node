package keeper

import (
	"context"
	"github.com/evmos/ethermint/x/committer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) LastProposalId(goCtx context.Context, req *types.QueryLastProposalIdRequest) (*types.QueryLastProposalIdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	lastProposal := k.GetLastProposal(ctx)
	return &types.QueryLastProposalIdResponse{
		LastProposalId: lastProposal.Id, 
		EndIndex: lastProposal.EndIndex}, nil
}

func (k Keeper) Proposal(goCtx context.Context, req *types.QueryProposalRequest) (*types.QueryProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	proposal, found := k.GetProposal(ctx, req.ProposalId)
	if !found {
		return &types.QueryProposalResponse{}, types.ErrNotExistProposal
	}
	return &types.QueryProposalResponse{Proposal: &proposal}, nil
}

