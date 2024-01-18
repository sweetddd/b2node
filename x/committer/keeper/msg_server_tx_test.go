package keeper_test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/evmos/ethermint/x/committer/types"
	"github.com/evmos/ethermint/x/committer/keeper"
	keepertest "github.com/evmos/ethermint/testutil/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestBatchProof(t *testing.T) {
	type tx struct {
		name 							string
		tx 								types.MsgBatchProofTx
		isError 					bool
		errMsg            string
		preRun 						func(ctx sdk.Context, k keeper.Keeper)
	}
	msgs := []tx{
		{
			name: "success",
			tx: types.MsgBatchProofTx{
				Id: 1,
				From: "from",
			},
			isError: false,
		},
		{
			name: "failed with incorrect proposal status",
			tx: types.MsgBatchProofTx{
				Id: 1,
				From: "from",
			},
			isError: true,
			errMsg: types.ErrProposalStatus.Error(),
			preRun: func(ctx sdk.Context, k keeper.Keeper) {
				k.SetProposal(ctx, types.Proposal{
					Id: 1,
					Status: types.Pending_Status,
					BlockHight: 10000,
				})
			},
		},
		{
			name: "failed with timeout proposal",
			tx: types.MsgBatchProofTx{
				Id: 1,
				From: "from",
			},
			isError: true,
			errMsg: types.ErrProposalTimeout.Error(),
			preRun: func(ctx sdk.Context, k keeper.Keeper) {
				k.SetProposal(ctx, types.Proposal{
					Id: 1,
					Status: types.Voting_Status,
					BlockHight: 1,
				})
			},
		},
	}

	for _, tc := range msgs {
		t.Run(tc.name, func(t *testing.T) {
			k, ctx := keepertest.CommitterKeeper(t)
			if tc.preRun != nil {
				tc.preRun(ctx, *k)
			}

			msgServer := keeper.NewMsgServerImpl(*k)

			_, err := msgServer.BatchProof(sdk.WrapSDKContext(ctx), &tc.tx)
			if tc.isError {
				require.Error(t, err)
				require.Equal(t, tc.errMsg, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestTapRoot(t *testing.T) {
	type tx struct {
		name 							string
		tx 								types.MsgTapRootTx
		isError 					bool
		errMsg            string
		preRun 						func(ctx sdk.Context, k keeper.Keeper)
	}
	msgs := []tx{
		{
			name: "success",
			tx: types.MsgTapRootTx{
				Id: 1,
				From: "from",
				BitcoinTxHash: "bitcoin_tx",
			},
			isError: false,
			preRun: func(ctx sdk.Context, k keeper.Keeper) {
				k.SetProposal(ctx, types.Proposal{
					Id: 1,
					Status: types.Pending_Status,
					Winner: "from",
					BlockHight: 10000,
				})
			},
		},	
		{
			name: "failed with not exist proposal",
			tx: types.MsgTapRootTx{
				Id: 1,
				From: "from",
			},
			isError: true,
			errMsg: types.ErrNotExistProposal.Error(),
		},	
		{
			name: "failed with incorrect account permission",
			tx: types.MsgTapRootTx{
				Id: 1,
				From: "from",
				BitcoinTxHash: "bitcoin_tx",
			},
			isError: true,
			errMsg: types.ErrAccountPermission.Error(),
			preRun: func(ctx sdk.Context, k keeper.Keeper) {
				k.SetProposal(ctx, types.Proposal{
					Id: 1,
					Status: types.Pending_Status,
					Winner: "winner",
					BlockHight: 10000,
				})
			},
		},
		{
			name: "failed with incorrect proposal status",
			tx: types.MsgTapRootTx{
				Id: 1,
				From: "from",
				BitcoinTxHash: "bitcoin_tx",
			},
			isError: true,
			errMsg: types.ErrProposalStatus.Error(),
			preRun: func(ctx sdk.Context, k keeper.Keeper) {
				k.SetProposal(ctx, types.Proposal{
					Id: 1,
					Winner: "from",
					Status: types.Succeed_Status,
					BlockHight: 10000,
				})
			},
		},
		{
			name: "failed with timeout proposal",
			tx: types.MsgTapRootTx{
				Id: 1,
				From: "from",
				BitcoinTxHash: "bitcoin_tx",
			},
			isError: true,
			errMsg: types.ErrProposalTimeout.Error(),
			preRun: func(ctx sdk.Context, k keeper.Keeper) {
				k.SetProposal(ctx, types.Proposal{
					Id: 1,
					Winner: "from",
					Status: types.Pending_Status,
					BlockHight: 1,
				})
			},
		},

	}

	for _, tc := range msgs {
		t.Run(tc.name, func(t *testing.T) {
			k, ctx := keepertest.CommitterKeeper(t)
			if tc.preRun != nil {
				tc.preRun(ctx, *k)
			}

			msgServer := keeper.NewMsgServerImpl(*k)

			_, err := msgServer.TapRoot(sdk.WrapSDKContext(ctx), &tc.tx)
			if tc.isError {
				require.Error(t, err)
				require.Equal(t, tc.errMsg, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestTimeoutProposal(t *testing.T) {
	type tx struct {
		name 							string
		tx 								types.MsgTimeoutProposalTx
		isError 					bool
		errMsg            string
		preRun 						func(ctx sdk.Context, k keeper.Keeper)
	}
	msgs := []tx{
		{
			name: "success",
			tx: types.MsgTimeoutProposalTx{
				Id: 1,
			},
			isError: false,
			preRun: func(ctx sdk.Context, k keeper.Keeper) {
				k.SetProposal(ctx, types.Proposal{
					Id: 1,
					Status: types.Pending_Status,
					BlockHight: 1,
				})
			},
		},
		{
			name: "failed with not exist proposal",
			tx: types.MsgTimeoutProposalTx{
				Id: 1,
			},
			isError: true,
			errMsg: types.ErrNotExistProposal.Error(),
		},
		{
			name: "failed with incorrect proposal status",
			tx: types.MsgTimeoutProposalTx{
				Id: 1,
			},
			isError: true,
			errMsg: types.ErrProposalStatus.Error(),
			preRun: func(ctx sdk.Context, k keeper.Keeper) {
				k.SetProposal(ctx, types.Proposal{
					Id: 1,
					Status: types.Succeed_Status,
					BlockHight: 1,
				})
			},
		},
		{
			name: "failed with not timeout proposal",
			tx: types.MsgTimeoutProposalTx{
				Id: 1,
			},
			isError: true,
			errMsg: types.ErrInvalidProposal.Error(),
			preRun: func(ctx sdk.Context, k keeper.Keeper) {
				k.SetProposal(ctx, types.Proposal{
					Id: 1,
					Status: types.Pending_Status,
					BlockHight: 10000,
				})
			},
		},
	}

	for _, tc := range msgs {
		t.Run(tc.name, func(t *testing.T) {
			k, ctx := keepertest.CommitterKeeper(t)
			if tc.preRun != nil {
				tc.preRun(ctx, *k)
			}

			msgServer := keeper.NewMsgServerImpl(*k)

			_, err := msgServer.TimeoutProposal(sdk.WrapSDKContext(ctx), &tc.tx)
			if tc.isError {
				require.Error(t, err)
				require.Equal(t, tc.errMsg, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
