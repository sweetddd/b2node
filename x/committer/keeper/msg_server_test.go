package keeper_test

import (
	"context"
	"testing"

  "github.com/evmos/ethermint/x/committer/types"
  "github.com/evmos/ethermint/x/committer/keeper"
	"github.com/evmos/ethermint/testutil"
	"github.com/stretchr/testify/require"
  keepertest "github.com/evmos/ethermint/testutil/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CommitterKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}


func TestSubmitProof(t *testing.T) {
	fromAddress := testutil.AccAddress()
	type tx struct {
		name 							string
		tx 								types.MsgSubmitProof
		isError 					bool
		errMsg            string
		preRun 						func(ctx sdk.Context, k keeper.Keeper)
	}
	msgs := []tx{
		{
			name: "success",
			tx: types.MsgSubmitProof{
				Id: 1,
				From: fromAddress,
			},
			isError: false,
		},
		{
			name: "failed with no permission",
			tx: types.MsgSubmitProof{
				Id: 1,
				From: "invalid_address",
			},
			isError: true,
			errMsg: types.ErrAccountPermission.Error(),
			preRun: func(ctx sdk.Context, k keeper.Keeper) {
				k.SetProposal(ctx, types.Proposal{
					Id: 1,
					Status: types.Voting_Status,
					BlockHight: 10000,
				})
			},
		},
		{
			name: "failed with incorrect proposal status",
			tx: types.MsgSubmitProof{
				Id: 1,
				From: fromAddress,
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
			tx: types.MsgSubmitProof{
				Id: 1,
				From: fromAddress,
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

			k.SetCommitter(ctx, types.Committer{
				CommitterList: []string{fromAddress},
			})

			msgServer := keeper.NewMsgServerImpl(*k)

			_, err := msgServer.SubmitProof(sdk.WrapSDKContext(ctx), &tc.tx)
			if tc.isError {
				require.Error(t, err)
				require.Equal(t, tc.errMsg, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestBitcoinTx(t *testing.T) {
	fromAddress := testutil.AccAddress()
	type tx struct {
		name 							string
		tx 								types.MsgBitcoinTx
		isError 					bool
		errMsg            string
		preRun 						func(ctx sdk.Context, k keeper.Keeper)
	}
	msgs := []tx{
		{
			name: "success",
			tx: types.MsgBitcoinTx{
				Id: 1,
				From: fromAddress	,
				BitcoinTxHash: "bitcoin_tx",
			},
			isError: false,
			preRun: func(ctx sdk.Context, k keeper.Keeper) {
				k.SetProposal(ctx, types.Proposal{
					Id: 1,
					Status: types.Pending_Status,
					Winner: fromAddress,
					BlockHight: 10000,
				})
			},
		},	
		{
			name: "failed with no permission",
			tx: types.MsgBitcoinTx{
				Id: 1,
				From: "invalid_address",
			},
			isError: true,
			errMsg: types.ErrAccountPermission.Error(),
			preRun: func(ctx sdk.Context, k keeper.Keeper) {
				k.SetProposal(ctx, types.Proposal{
					Id: 1,
					Status: types.Pending_Status,
					Winner: fromAddress,
					BlockHight: 10000,
				})
			},
		},
		{
			name: "failed with not exist proposal",
			tx: types.MsgBitcoinTx{
				Id: 1,
				From: fromAddress,
			},
			isError: true,
			errMsg: types.ErrNotExistProposal.Error(),
		},	
		{
			name: "failed with incorrect account permission",
			tx: types.MsgBitcoinTx{
				Id: 1,
				From: fromAddress,
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
			tx: types.MsgBitcoinTx{
				Id: 1,
				From: fromAddress,
				BitcoinTxHash: "bitcoin_tx",
			},
			isError: true,
			errMsg: types.ErrProposalStatus.Error(),
			preRun: func(ctx sdk.Context, k keeper.Keeper) {
				k.SetProposal(ctx, types.Proposal{
					Id: 1,
					Winner: fromAddress,
					Status: types.Succeed_Status,
					BlockHight: 10000,
				})
			},
		},
		{
			name: "failed with timeout proposal",
			tx: types.MsgBitcoinTx{
				Id: 1,
				From: fromAddress,
				BitcoinTxHash: "bitcoin_tx",
			},
			isError: true,
			errMsg: types.ErrProposalTimeout.Error(),
			preRun: func(ctx sdk.Context, k keeper.Keeper) {
				k.SetProposal(ctx, types.Proposal{
					Id: 1,
					Winner: fromAddress,
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

			k.SetCommitter(ctx, types.Committer{
				CommitterList: []string{fromAddress},
			})

			msgServer := keeper.NewMsgServerImpl(*k)

			_, err := msgServer.BitcoinTx(sdk.WrapSDKContext(ctx), &tc.tx)
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
	fromAddress := testutil.AccAddress()
	type tx struct {
		name 							string
		tx 								types.MsgTimeoutProposal
		isError 					bool
		errMsg            string
		preRun 						func(ctx sdk.Context, k keeper.Keeper)
	}
	msgs := []tx{
		{
			name: "success",
			tx: types.MsgTimeoutProposal{
				Id: 1,
				From: fromAddress,
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
			name: "failed with no permission",
			tx: types.MsgTimeoutProposal{
				Id: 1,
				From: "invalid_address",
			},
			isError: true,
			errMsg: types.ErrAccountPermission.Error(),
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
			tx: types.MsgTimeoutProposal{
				Id: 1,
				From: fromAddress,
			},
			isError: true,
			errMsg: types.ErrNotExistProposal.Error(),
		},
		{
			name: "failed with incorrect proposal status",
			tx: types.MsgTimeoutProposal{
				Id: 1,
				From: fromAddress,
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
			tx: types.MsgTimeoutProposal{
				Id: 1,
				From: fromAddress,
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

			k.SetCommitter(ctx, types.Committer{
				CommitterList: []string{fromAddress},
			})

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

func TestAddCommitter(t *testing.T) {
	fromAddress := testutil.AccAddress()
	type tx struct {
		name 							string
		tx 								types.MsgAddCommitter
		isError 					bool
		errMsg            string
		preRun 						func(ctx sdk.Context, k keeper.Keeper)
	}

	msgs := []tx{
		{
			name: "success",
			tx: types.MsgAddCommitter{
				From: fromAddress,
				Committer: fromAddress,
			},
			isError: false,
		},
		{
			name: "failed with no permission",
			tx: types.MsgAddCommitter{
				From: "invalid_address",
				Committer: fromAddress,
			},
			isError: true,
			errMsg: types.ErrAccountPermission.Error(),
		},
		{
			name: "failed with already exist committer",
			tx: types.MsgAddCommitter{
				From: fromAddress,
				Committer: fromAddress,
			},
			isError: true,
			errMsg: types.ErrExistCommitter.Error(),
			preRun: func(ctx sdk.Context, k keeper.Keeper) {
				k.SetCommitter(ctx, types.Committer{
					CommitterList: []string{fromAddress},
				})
			},
		},
	}

	for _, tc := range msgs {
		t.Run(tc.name, func(t *testing.T) {
			k, ctx := keepertest.CommitterKeeper(t)

			adminPolicy := types.AdminPolicy{
				Address: fromAddress,
				PolicyType: types.PolicyType_group1,
			}
			params := types.Params{
				AdminPolicy: []*types.AdminPolicy{&adminPolicy},
			}

			k.SetParams(ctx, params)

			if tc.preRun != nil {
				tc.preRun(ctx, *k)
			}

			msgServer := keeper.NewMsgServerImpl(*k)

			_, err := msgServer.AddCommitter(sdk.WrapSDKContext(ctx), &tc.tx)
			if tc.isError {
				require.Error(t, err)
				require.Equal(t, tc.errMsg, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestRemoveCommitter(t *testing.T) {
	addr := testutil.AccAddress()
	addr_no_permission := testutil.AccAddress()
	type tx struct {
		name 							string
		msg 								types.MsgRemoveCommitter
		isError 					bool
		errMsg            string
		preRun 						func(ctx sdk.Context, k keeper.Keeper)
	}

	msgs := []tx{
		{
			name: "success",
			msg: types.MsgRemoveCommitter{
				From: addr,
				Committer: addr,
			},
			isError: false,
			preRun: func(ctx sdk.Context, k keeper.Keeper) {
				k.SetCommitter(ctx, types.Committer{
					CommitterList: []string{addr},
				})
			},
		},
		{
			name: "failed with no permission",
			msg: types.MsgRemoveCommitter{
				From: addr_no_permission,
				Committer: addr,
			},
			isError: true,
			errMsg: types.ErrAccountPermission.Error(),
		},
		{
			name: "failed with not exist committer",
			msg: types.MsgRemoveCommitter{
				From: addr,
				Committer: addr,
			},
			isError: true,
			errMsg: types.ErrNotExistCommitter.Error(),
		},
	}

	for _, tc := range msgs {
		t.Run(tc.name, func(t *testing.T) {
			k, ctx := keepertest.CommitterKeeper(t)

			adminPolicy := types.AdminPolicy{
				Address: addr,
				PolicyType: types.PolicyType_group1,
			}
			params := types.Params{
				AdminPolicy: []*types.AdminPolicy{&adminPolicy},
			}

			k.SetParams(ctx, params)

			if tc.preRun != nil {
				tc.preRun(ctx, *k)
			}

			msgServer := keeper.NewMsgServerImpl(*k)

			_, err := msgServer.RemoveCommitter(sdk.WrapSDKContext(ctx), &tc.msg)
			if tc.isError {
				require.Error(t, err)
				require.Equal(t, tc.errMsg, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
