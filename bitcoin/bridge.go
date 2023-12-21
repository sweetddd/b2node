package bitcoin

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"net/url"
	"os"
	"path"

	b2aa "github.com/b2network/b2-go-aa-utils"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Bridge bridge
// TODO: only L1 -> L2, More calls may be supported later
type Bridge struct {
	EthRPCURL       string
	EthPrivKey      *ecdsa.PrivateKey
	ContractAddress common.Address
	ABI             string
	GasLimit        uint64
	// AA contract address
	AASCARegistry   common.Address
	AAKernelFactory common.Address
}

// NewBridge new bridge
func NewBridge(bridgeCfg BridgeConfig, abiFileDir string) (*Bridge, error) {
	rpcURL, err := url.ParseRequestURI(bridgeCfg.EthRPCURL)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(bridgeCfg.EthPrivKey)
	if err != nil {
		return nil, err
	}

	abi, err := os.ReadFile(path.Join(abiFileDir, bridgeCfg.ABI))
	if err != nil {
		return nil, err
	}

	return &Bridge{
		EthRPCURL:       rpcURL.String(),
		ContractAddress: common.HexToAddress(bridgeCfg.ContractAddress),
		EthPrivKey:      privateKey,
		ABI:             string(abi),
		GasLimit:        bridgeCfg.GasLimit,
		AASCARegistry:   common.HexToAddress(bridgeCfg.AASCARegistry),
		AAKernelFactory: common.HexToAddress(bridgeCfg.AAKernelFactory),
	}, nil
}

// Deposit to ethereum
func (b *Bridge) Deposit(hash string, bitcoinAddress string, amount int64) ([]byte, error) {
	if bitcoinAddress == "" {
		return nil, fmt.Errorf("bitcoin address is empty")
	}

	if hash == "" {
		return nil, fmt.Errorf("tx id is empty")
	}

	ctx := context.Background()

	toAddress, err := b.BitcoinAddressToEthAddress(bitcoinAddress)
	if err != nil {
		return nil, fmt.Errorf("btc address to eth address err:%w", err)
	}

	txHash, err := chainhash.NewHashFromStr(hash)
	if err != nil {
		return nil, err
	}

	data, err := b.ABIPack(b.ABI, "depositV2", txHash, common.HexToAddress(toAddress), new(big.Int).SetInt64(amount))
	if err != nil {
		return nil, fmt.Errorf("abi pack err:%w", err)
	}

	return b.ethContractCall(ctx, b.EthPrivKey, data)
}

func (b *Bridge) ethContractCall(ctx context.Context, priv *ecdsa.PrivateKey, data []byte) ([]byte, error) {
	client, err := ethclient.Dial(b.EthRPCURL)
	if err != nil {
		return nil, err
	}

	publicKey := priv.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	callMsg := ethereum.CallMsg{
		From: crypto.PubkeyToAddress(*publicKeyECDSA),
		To:   &b.ContractAddress,
		Data: data,
	}

	return client.CallContract(ctx, callMsg, nil)
}

// ABIPack the given method name to conform the ABI. Method call's data
func (b *Bridge) ABIPack(abiData string, method string, args ...interface{}) ([]byte, error) {
	contractAbi, err := abi.JSON(bytes.NewReader([]byte(abiData)))
	if err != nil {
		return nil, err
	}
	return contractAbi.Pack(method, args...)
}

// BitcoinAddressToEthAddress bitcoin address to eth address
func (b *Bridge) BitcoinAddressToEthAddress(bitcoinAddress string) (string, error) {
	client, err := ethclient.Dial(b.EthRPCURL)
	if err != nil {
		return "", err
	}

	targetEthAddress, err := b2aa.GetSCAAddress(client, b.AASCARegistry, b.AAKernelFactory, bitcoinAddress)
	if err != nil {
		return "", err
	}
	return targetEthAddress.String(), nil
}
