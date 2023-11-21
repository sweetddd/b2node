package keeper

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type BITCOINConfig struct {
	// NetworkName defines the bitcoin network name
	NetworkName string `yaml:"network-name"`
	RPCHost     string `yaml:"rpc-host"`
	RPCPort     string `yaml:"rpc-port"`
	RPCUser     string `yaml:"rpc-user"`
	RPCPass     string `yaml:"rpc-pass"`
	WalletName  string `yaml:"wallet-name"`
	// Destination defines the taproot transaction destination address
	Destination string `yaml:"destination"`
}

const (
	BitcoinRPCConfigFileName = "bitcoin.yaml"
)

func DefaultBITCOINConfig(homePath string) *BITCOINConfig {
	_, err := os.Stat(homePath + "/" + BitcoinRPCConfigFileName)
	config := BITCOINConfig{}
	if os.IsNotExist(err) {
		config.NetworkName = "signet"
		config.RPCHost = "localhost"
		config.RPCPort = "8332"
		config.RPCUser = "rpcuser"
		config.RPCPass = "rpcpass"
		config.WalletName = "walletname"
		config.Destination = "tb1qgm39cu009lyvq93afx47pp4h9wxq5x92lxxgnz"
		return &config
	}

	data, err := os.ReadFile(homePath + "/" + BitcoinRPCConfigFileName)
	if err != nil {
		log.Fatalf("can not read rpc config file: %v", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("can not unmarshal rpc config file: %v", err)
	}
	return &config
}
