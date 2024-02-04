package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/testutil"
	testkeeper "github.com/evmos/ethermint/testutil/keeper"
	"github.com/evmos/ethermint/x/committer/types"
	"github.com/stretchr/testify/require"
)

func TestGetLastProposal(t *testing.T) {
	keeper, ctx := testkeeper.CommitterKeeper(t)

	keeper.SetLastProposal(ctx, types.Proposal{
		Id:       1,
		EndIndex: 1,
	})

	response, err := keeper.LastProposalID(sdk.WrapSDKContext(ctx), &types.QueryLastProposalIdRequest{})
	require.NoError(t, err)

	require.Equal(t, &types.QueryLastProposalIdResponse{
		LastProposalId: 1,
		EndIndex:       1,
	}, response)
}

func TestGetProposal(t *testing.T) {
	keeper, ctx := testkeeper.CommitterKeeper(t)

	keeper.SetProposal(ctx, types.Proposal{
		Id: 1,
	})

	response, err := keeper.Proposal(sdk.WrapSDKContext(ctx), &types.QueryProposalRequest{ProposalId: 1})
	require.NoError(t, err)

	require.Equal(t, &types.QueryProposalResponse{
		Proposal: &types.Proposal{
			Id: 1,
		},
	}, response)
}

func TestGetCommitters(t *testing.T) {
	addr1 := testutil.AccAddress()
	addr2 := testutil.AccAddress()
	keeper, ctx := testkeeper.CommitterKeeper(t)

	committers := types.Committer{
		CommitterList: []string{addr1, addr2},
	}

	keeper.SetCommitter(ctx, committers)

	response, err := keeper.Committers(sdk.WrapSDKContext(ctx), &types.QueryCommitterRequest{})
	require.NoError(t, err)

	require.Equal(t, &types.QueryCommitterResponse{
		Committers: &committers,
	}, response)
}
