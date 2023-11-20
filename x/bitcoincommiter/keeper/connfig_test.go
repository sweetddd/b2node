package keeper_test

import (
	"testing"

	"github.com/evmos/ethermint/x/bitcoincommiter/keeper"
	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	config := keeper.DefaultBITCOINConfig("../test/source")
	require.Equal(t, "signet", config.NetworkName)
	require.Equal(t, "localhost", config.RPCHost)
	require.Equal(t, "8332", config.RPCPort)
	require.Equal(t, "b2node", config.RPCUser)
	require.Equal(t, "b2node", config.RPCPass)
	require.Equal(t, "b2node", config.WalletName)
	require.Equal(t, "tb1qgm39cu009lyvq93afx47pp4h9wxq5x92lxxgnz", config.Destination)
}
