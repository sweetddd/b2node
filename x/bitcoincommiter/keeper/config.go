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
	data, err := os.ReadFile(homePath + "/" + BitcoinRPCConfigFileName)
	if err != nil {
		log.Fatalf("can not read rpc config file: %v", err)
	}
	config := BITCOINConfig{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("can not unmarshal rpc config file: %v", err)
	}
	return &config
}
