package bitcoin_test

import (
	"errors"
	"log"
	"testing"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/evmos/ethermint/bitcoin"
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
	ethlient, err := ethclient.Dial(bitcoinCfg.Bridge.EthRPCURL)
	if err != nil {
		log.Fatalf("EVMListenerService failed to create eth client: %v", err)
	}
	defer func() {
		ethlient.Close()
	}()

	listenerService := bitcoin.NewEVMListenerService(btclient, ethlient, bitcoinCfg)
	require.NotNil(t, listenerService)
}

func TestIsUnspentTX(t *testing.T) {
	listenerService := evmListenerWithConfig(t)

	testCase := []struct {
		name  string
		value bool
	}{
		{
			name:  "NoUnspentTX",
			value: false,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			isOK, _ := listenerService.IsUnspentTX()
			require.Equal(t, tc.value, isOK)
		})
	}
}

func TestTransferToBtc(t *testing.T) {
	listenerService := evmListenerWithConfig(t)

	testCase := []struct {
		name      string
		addresses []string
		amounts   []int64
		err       error
	}{
		{
			name:      "fail",
			addresses: []string{"tb1qgm39cu009lyvq93afx47pp4h9wxq5x92lxxgnz"},
			amounts:   []int64{10000},
			err:       errors.New("the client has been shutdown"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			err := listenerService.TransferToBtc(tc.addresses, tc.amounts)
			if err != nil {
				require.Error(t, tc.err, err)
			}
		})
	}
}

func evmListenerWithConfig(t *testing.T) *bitcoin.EVMListenerService {
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

	// start eth rpc client
	ethlient, err := ethclient.Dial(bitcoinCfg.Bridge.EthRPCURL)
	if err != nil {
		log.Fatalf("EVMListenerService failed to create eth client: %v", err)
	}
	defer func() {
		ethlient.Close()
	}()

	listenerService := bitcoin.NewEVMListenerService(btclient, ethlient, bitcoinCfg)
	return listenerService
}
