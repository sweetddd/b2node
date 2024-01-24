package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	ethermint "github.com/evmos/ethermint/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

func NewAccountKeeper(t testing.TB) (*keeper.AccountKeeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(authtypes.StoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(
		cdc,
		nil,
		storeKey,
		nil,
		"account",
	)

	keeper := keeper.NewAccountKeeper(
		cdc,
		storeKey,
		paramsSubspace,
		ethermint.ProtoAccount,
		nil,
		sdk.GetConfig().GetBech32AccountAddrPrefix(),
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{Height: 10002}, false, log.NewNopLogger())

	return &keeper, ctx
}
