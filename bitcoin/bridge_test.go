package bitcoin_test

import (
	"os"
	"path"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/evmos/ethermint/bitcoin"
	"github.com/stretchr/testify/assert"
)

func TestNewBridge(t *testing.T) {
	abiPath := path.Join("./testdata")

	abi, err := os.ReadFile(path.Join("./testdata", "abi.json"))
	if err != nil {
		t.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef")
	if err != nil {
		t.Fatal(err)
	}

	bridgeCfg := bitcoin.BridgeConfig{
		EthRPCURL:       "http://localhost:8545",
		ContractAddress: "0x123456789abcdef",
		EthPrivKey:      "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
		ABI:             "abi.json",
		GasLimit:        1000000,
	}

	bridge, err := bitcoin.NewBridge(bridgeCfg, abiPath)
	assert.NoError(t, err)
	assert.NotNil(t, bridge)
	assert.Equal(t, bridgeCfg.EthRPCURL, bridge.EthRPCURL)
	assert.Equal(t, common.HexToAddress("0x123456789abcdef"), bridge.ContractAddress)
	assert.Equal(t, privateKey, bridge.EthPrivKey)
	assert.Equal(t, string(abi), bridge.ABI)
	assert.Equal(t, bridgeCfg.GasLimit, bridge.GasLimit)
}
