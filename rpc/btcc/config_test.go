package btcc

import (
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/evmos/ethermint/server/config"
	"github.com/stretchr/testify/require"
)

// TestSetBitcoinConfig tests the SetBitcoinConfig function.
func TestSetBitcoinConfig(t *testing.T) {
	// Create a new BitcoinConfig instance.
	config := config.BITCOINConfig{
		NetworkName: "mainnet",
		RPCHost:     "localhost",
		RPCPort:     "8332",
		RPCUser:     "b2node",
		RPCPass:     "b2node",
		WalletName:  "b2node",
	}

	// Call the SetBitcoinConfig function.
	rpcConfig := SetBitcoinConfig(config)
	require.Equal(t, rpcConfig.Params, chaincfg.MainNetParams)
	require.Equal(t, rpcConfig.ConnConfig.Host, config.RPCHost+":"+config.RPCPort+"/wallet/"+config.WalletName)
	require.Equal(t, rpcConfig.ConnConfig.User, config.RPCUser)
	require.Equal(t, rpcConfig.ConnConfig.Pass, config.RPCPass)
}
