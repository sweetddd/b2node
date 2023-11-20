package keeper

import (
	"testing"

	codecc "github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	bitcoincommitertypes "github.com/evmos/ethermint/x/bitcoincommiter/types"
	"github.com/stretchr/testify/require"
)

func TestInitKeeper(t *testing.T) {
	amino := codecc.NewLegacyAmino()
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	marshaler := codecc.NewProtoCodec(interfaceRegistry)
	key := sdk.NewKVStoreKey("key")
	tkey := sdk.NewTransientStoreKey("tkey")
	paramsKeeper := paramskeeper.NewKeeper(marshaler, amino, key, tkey)
	space := paramsKeeper.Subspace(bitcoincommitertypes.ModuleName)
	commiterKey := sdk.NewKVStoreKey(bitcoincommitertypes.ModuleName)
	keeper := NewKeeper(marshaler, commiterKey, space, "../test/source")
	require.Equal(t, "signet", keeper.config.NetworkName)
	require.Equal(t, "localhost", keeper.config.RPCHost)
	require.Equal(t, "8332", keeper.config.RPCPort)
	require.Equal(t, "b2node", keeper.config.RPCUser)
	require.Equal(t, "b2node", keeper.config.RPCPass)
	require.Equal(t, "b2node", keeper.config.WalletName)
	require.Equal(t, "tb1qgm39cu009lyvq93afx47pp4h9wxq5x92lxxgnz", keeper.config.Destination)
}
