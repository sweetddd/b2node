// Copyright 2021 Evmos Foundation
// This file is part of Evmos' Ethermint library.
//
// The Ethermint library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Ethermint library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Ethermint library. If not, see https://github.com/evmos/ethermint/blob/main/LICENSE
package bitcoin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	btcrpcclient "github.com/btcsuite/btcd/rpcclient"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/tendermint/tendermint/libs/service"
)

const (
	ListenerServiceName = "EVMListenerService"
)

// EVMListenerService indexes transactions for json-rpc service.
type EVMListenerService struct {
	service.BaseService

	btcCli *btcrpcclient.Client
	ethCli *ethclient.Client
	config *BitconConfig
}

// NewEVMListenerService returns a new service instance.
func NewEVMListenerService(
	btcCli *btcrpcclient.Client,
	ethCli *ethclient.Client,
	config *BitconConfig,
) *EVMListenerService {
	is := &EVMListenerService{btcCli: btcCli, ethCli: ethCli, config: config}
	is.BaseService = *service.NewBaseService(nil, ListenerServiceName, is)
	return is
}

// OnStart implements service.Service by subscribing for new blocks
// and indexing them by events.
func (eis *EVMListenerService) OnStart() error {
	lastBlock := eis.config.Evm.StartHeight
	addresses := []common.Address{
		common.HexToAddress(eis.config.Evm.ContractAddress),
	}
	topics := [][]common.Hash{
		{
			common.HexToHash(eis.config.Evm.Deposit),
			common.HexToHash(eis.config.Evm.Withdraw),
		},
	}
	for {
		height, err := eis.ethCli.BlockNumber(context.Background())
		if err != nil {
			eis.Logger.Error("EVMListenerService HeaderByNumber is failed:", "err", err)
			time.Sleep(time.Second * 10)
			continue
		}

		latestBlock, err := strconv.ParseInt(fmt.Sprint(height), 10, 64)
		if err != nil {
			eis.Logger.Error("EVMListenerService ParseInt latestBlock", "err", err)
			return err
		}
		eis.Logger.Info("EVMListenerService ethClient height", "height", latestBlock)

		if latestBlock <= lastBlock {
			time.Sleep(time.Second * 10)
			continue
		}

		for i := lastBlock + 1; i <= latestBlock; i++ {
			query := ethereum.FilterQuery{
				FromBlock: big.NewInt(i),
				ToBlock:   big.NewInt(i),
				Topics:    topics,
				Addresses: addresses,
			}
			logs, err := eis.ethCli.FilterLogs(context.Background(), query)
			if err != nil {
				eis.Logger.Error("EVMListenerService failed to fetch block", "height", i, "err", err)
				break
			}
			for _, vlog := range logs {
				eventHash := common.BytesToHash(vlog.Topics[0].Bytes())
				if eventHash == common.HexToHash(eis.config.Evm.Deposit) {
					// todo
					data := DepositEvent{
						Sender:    TopicToAddress(vlog, 1),
						ToAddress: TopicToAddress(vlog, 2),
						Amount:    DataToBigInt(vlog, 0),
					}
					value, err := json.Marshal(&data)
					if err != nil {
						eis.Logger.Error("EVMListenerService listener deposit Marshal failed: ", "err", err)
						return err
					}
					eis.Logger.Info("EVMListenerService listener deposit event: ", "num", i, "deposit", string(value))
				} else if eventHash == common.HexToHash(eis.config.Evm.Withdraw) {
					data := WithdrawEvent{
						FromAddress: TopicToAddress(vlog, 1),
						ToAddress:   DataToString(vlog, 0),
						Amount:      DataToBigInt(vlog, 1),
					}
					value, err := json.Marshal(&data)
					if err != nil {
						eis.Logger.Error("EVMListenerService listener withdraw Marshal failed: ", "err", err)
						return err
					}
					eis.Logger.Info("EVMListenerService listener withdraw event: ", "num", i, "withdraw", string(value))

					amount := DataToBigInt(vlog, 1)
					err = eis.transferToBtc(DataToString(vlog, 0), amount.Int64())
					if err != nil {
						eis.Logger.Error("EVMListenerService transferToBtc failed: ", "err", err)
						// deal with Insufficient balance
						if err.Error() == "unable to calculate change amount" {
							time.Sleep(5 * time.Minute)
							continue
						}
						time.Sleep(1 * time.Minute)
						continue
					}
					// eis.Logger.Info("EVMListenerService listener withdraw event: ", "withdraw", string(value))
				}
			}
			lastBlock = i
		}
	}
}

func (eis *EVMListenerService) transferToBtc(destAddrStr string, amount int64) error {
	eis.Logger.Info("EVMListenerService btc transfer", "destAddrStr", destAddrStr, "amount", amount)
	sourceAddrStr := eis.config.SourceAddress

	var defaultNet *chaincfg.Params
	networkName := eis.config.NetworkName
	defaultNet = ChainParams(networkName)

	// get sourceAddress UTXO
	sourceAddr, err := btcutil.DecodeAddress(sourceAddrStr, defaultNet)
	if err != nil {
		eis.Logger.Error("EVMListenerService transferToBtc DecodeAddress failed: ", "err", err)
		return err
	}

	unspentTxs, err := eis.btcCli.ListUnspentMinMaxAddresses(1, 9999999, []btcutil.Address{sourceAddr})
	if err != nil {
		eis.Logger.Error("EVMListenerService ListUnspentMinMaxAddresses transferToBtc DecodeAddress failed: ", "err", err)
		return err
	}

	inputs := make([]btcjson.TransactionInput, 0, 10)
	totalInputAmount := btcutil.Amount(0)
	for _, unspentTx := range unspentTxs {
		inputs = append(inputs, btcjson.TransactionInput{
			Txid: unspentTx.TxID,
			Vout: unspentTx.Vout,
		})
		totalInputAmount += btcutil.Amount(unspentTx.Amount * 1e8)
		if (int64(totalInputAmount) + eis.config.Fee) > amount {
			break
		}
	}
	// eis.Logger.Info("ListUnspentMinMaxAddresses", "totalInputAmount", totalInputAmount)
	changeAmount := int64(totalInputAmount) - eis.config.Fee - amount // fee
	if changeAmount > 0 {
		changeAddr, err := btcutil.DecodeAddress(sourceAddrStr, defaultNet)
		if err != nil {
			eis.Logger.Error("EVMListenerService transferToBtc DecodeAddress sourceAddress failed: ", "err", err)
			return err
		}
		destAddr, err := btcutil.DecodeAddress(destAddrStr, defaultNet)
		if err != nil {
			eis.Logger.Error("EVMListenerService transferToBtc DecodeAddress destAddress failed: ", "err", err)
			return err
		}
		outputs := map[btcutil.Address]btcutil.Amount{
			changeAddr: btcutil.Amount(changeAmount),
			destAddr:   btcutil.Amount(amount),
		}
		rawTx, err := eis.btcCli.CreateRawTransaction(inputs, outputs, nil)
		if err != nil {
			eis.Logger.Error("EVMListenerService transferToBtc CreateRawTransaction failed: ", "err", err)
			return err
		}

		// sign
		signedTx, complete, err := eis.btcCli.SignRawTransactionWithWallet(rawTx)
		if err != nil {
			eis.Logger.Error("EVMListenerService transferToBtc SignRawTransactionWithWallet failed: ", "err", err)
			return err
		}
		if !complete {
			eis.Logger.Error("EVMListenerService transferToBtc SignRawTransactionWithWallet failed: ", "err", errors.New("SignRawTransaction not complete"))
			return errors.New("SignRawTransaction not complete")
		}
		// send
		txHash, err := eis.btcCli.SendRawTransaction(signedTx, true)
		if err != nil {
			eis.Logger.Error("EVMListenerService transferToBtc SendRawTransaction failed: ", "err", err)
			return err
		}
		eis.Logger.Info("EVMListenerService tx success: ", "from", sourceAddrStr, "to", destAddrStr, "amount", amount, "hash", txHash.String())
		return nil
	}

	return errors.New("unable to calculate change amount")
}
