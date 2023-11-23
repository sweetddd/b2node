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

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Bridge struct {
	EthRPCURL       string
	EthPrivKey      *ecdsa.PrivateKey
	ContractAddress common.Address
	ABI             string
	GasLimit        uint64
}

// NewBridge new bridge
func NewBridge(bridgeCfg BridgeConfig, abiFileDir string) (*Bridge, error) {
	rpcUrl, err := url.ParseRequestURI(bridgeCfg.EthRPCURL)
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
		EthRPCURL:       rpcUrl.String(),
		ContractAddress: common.HexToAddress(bridgeCfg.ContractAddress),
		EthPrivKey:      privateKey,
		ABI:             string(abi),
		GasLimit:        bridgeCfg.GasLimit,
	}, nil
}

// Deposit to ethereum
// TODO:  partition method and add test
func (b *Bridge) Deposit(bitcoinAddress string, amount int64) error {
	if bitcoinAddress == "" {
		return fmt.Errorf("bitcoin address is empty")
	}

	ctx := context.Background()
	// dail ethereum rpc
	client, err := ethclient.Dial(b.EthRPCURL)
	if err != nil {
		return err
	}

	publicKey := b.EthPrivKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("error casting public key to ECDSA")
	}
	nonce, err := client.PendingNonceAt(ctx, crypto.PubkeyToAddress(*publicKeyECDSA))
	if err != nil {
		return err
	}
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}

	contractAbi, err := abi.JSON(bytes.NewReader([]byte(b.ABI)))
	if err != nil {
		return err
	}

	toAddress, err := b.BitcoinAddressToEthAddress(bitcoinAddress)
	if err != nil {
		return err
	}

	data, err := contractAbi.Pack("deposit", common.HexToAddress(toAddress), new(big.Int).SetInt64(amount))
	if err != nil {
		return err
	}

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &b.ContractAddress,
		Value:    big.NewInt(0),
		Gas:      b.GasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return err
	}
	// sign tx
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), b.EthPrivKey)
	if err != nil {
		return err
	}

	// send tx
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return err
	}

	// wait tx confirm
	receipt, err := bind.WaitMined(ctx, client, signedTx)
	if err != nil {
		return err
	}

	if receipt.Status != 1 {
		return fmt.Errorf("tx failed, receipt:%v", receipt)
	}

	return nil
}

// BitcoinAddressToEthAddress bitcoin address to eth address
// TODO: implementation
func (b *Bridge) BitcoinAddressToEthAddress(bitcoinAddress string) (string, error) {
	// TODO: wait aa finished
	return bitcoinAddress, nil
}
