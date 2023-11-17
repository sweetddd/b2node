package bitcoin

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/evmos/ethermint/server/config"
)

// Config this param use for holding bitcoin rpc config
var Config RpcConfig

// RpcConfig this struct use for storing bitcoin rpc config
type RpcConfig struct {
	ConnConfig rpcclient.ConnConfig
	Params     chaincfg.Params
}

// SetBitcoinConfig this method uses for setting bitcoin rpc
func SetBitcoinConfig(config config.BITCOINConfig) RpcConfig {
	var params chaincfg.Params
	switch config.NetworkName {
	case chaincfg.SigNetParams.Name:
		params = chaincfg.SigNetParams
	case chaincfg.MainNetParams.Name:
		params = chaincfg.MainNetParams
	case chaincfg.TestNet3Params.Name:
		params = chaincfg.TestNet3Params
	case chaincfg.SimNetParams.Name:
		params = chaincfg.SimNetParams
	case chaincfg.RegressionNetParams.Name:
		params = chaincfg.RegressionNetParams
	default:
		params = chaincfg.MainNetParams
	}
	Config = RpcConfig{
		ConnConfig: rpcclient.ConnConfig{
			Host: config.RpcHost,
		},
		Params: params,
	}
	return Config
}
