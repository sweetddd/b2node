package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/evmos/ethermint/x/committer/keeper"
	"github.com/evmos/ethermint/x/committer/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

func setupKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"b2-node",
	)

	k := keeper.NewKeeper(cdc, storeKey, memStoreKey, paramsSubspace)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return k, ctx
}

func TestSetProposal(t *testing.T) {
	proposal := types.Proposal{
		Id:                    1,
		Proposer:              "proposer",
		ProofHash:             "proof_hash",
		StateRootHash:         "state_root_hash",
		StartIndex:            1,
		EndIndex:              2,
		BlockHight:            1,
		Status:                1,
		BitcoinTxHash:         "bitcoin_tx",
		Winner:                "winner",
		VotedListPhaseCommit:  []string{"voted_list_phase_commit"},
		VotedListPhaseTimeout: []string{"voted_list_phase_timeout"},
	}

	k, ctx := setupKeeper(t)
	k.SetProposal(ctx, proposal)

	p, found := k.GetProposal(ctx, proposal.Id)
	require.True(t, found)
	require.Equal(t, proposal, p)

	k.SetLastProposal(ctx, proposal)

	lastProposal := k.GetLastProposal(ctx)
	require.Equal(t, proposal.Id, lastProposal.Id)
}

func TestSetCommitter(t *testing.T) {
	k, ctx := setupKeeper(t)

	committers := types.Committer{
		CommitterList: []string{"committer1, committer2"},
	}

	k.SetCommitter(ctx, committers)

	c := k.GetAllCommitters(ctx)
	require.Equal(t, committers, c)
}
