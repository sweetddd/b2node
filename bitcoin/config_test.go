package bitcoin

import (
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/evmos/ethermint/server/config"
	"github.com/stretchr/testify/require"
)

// TestSetBitcoinConfig tests the SetBitcoinConfig function.
func TestSetBitcoinConfig(t *testing.T) {
	// Create a new BitcoinConfig instance.
	config := &config.BITCOINConfig{
		NetworkName: "mainnet",
		RpcHost:     "username",
		RpcUser:     "password",
		RpcPass:     "localhost",
		WalletName:  "ss",
	}

	// Call the SetBitcoinConfig function.
	rpcConfig := SetBitcoinConfig(config)
	require.Equal(t, rpcConfig.Params, chaincfg.MainNetParams)
	require.Equal(t, rpcConfig.ConnConfig.Host, config.RpcHost)
	require.Equal(t, rpcConfig.ConnConfig.User, config.RpcUser)
	require.Equal(t, rpcConfig.ConnConfig.Pass, config.RpcPass)
}
