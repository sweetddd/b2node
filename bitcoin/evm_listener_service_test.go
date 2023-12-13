package bitcoin_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/evmos/ethermint/bitcoin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEVMListenerService(t *testing.T) {
	bitcoinCfg, err := bitcoin.LoadBitcoinConfig("./testdata")
	require.NoError(t, err)

	// start btc rpc client
	btclient, err := rpcclient.New(&rpcclient.ConnConfig{
		Host:         bitcoinCfg.RPCHost + ":" + bitcoinCfg.RPCPort,
		User:         bitcoinCfg.RPCUser,
		Pass:         bitcoinCfg.RPCPass,
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}, nil)
	if err != nil {
		log.Fatalf("EVMListenerService failed to create bitcoin client: %v", err)
	}
	defer func() {
		btclient.Shutdown()
	}()

	// start eth rpc client
	ethlient, err := ethclient.Dial(fmt.Sprintf("%s:%s", bitcoinCfg.Evm.RPCHost, bitcoinCfg.Evm.RPCPort))
	if err != nil {
		log.Fatalf("EVMListenerService failed to create eth client: %v", err)
	}
	defer func() {
		ethlient.Close()
	}()

	listenerService := bitcoin.NewEVMListenerService(btclient, ethlient, bitcoinCfg)
	assert.NotNil(t, listenerService)
}
