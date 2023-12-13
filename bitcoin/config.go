package bitcoin

import (
	"os"
	"path"
	"strings"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/spf13/viper"
)

// BitconConfig defines the bitcoin config
// TODO: defined different config group eg: bitcoin, bridge, indexer, commiter
type BitconConfig struct {
	// NetworkName defines the bitcoin network name
	NetworkName string `mapstructure:"network-name"`
	// RPCHost defines the bitcoin rpc host
	RPCHost string `mapstructure:"rpc-host"`
	// RPCPort defines the bitcoin rpc port
	RPCPort string `mapstructure:"rpc-port"`
	// RPCUser defines the bitcoin rpc user
	RPCUser string `mapstructure:"rpc-user"`
	// RPCPass defines the bitcoin rpc password
	RPCPass string `mapstructure:"rpc-pass"`
	// WalletName defines the bitcoin wallet name
	WalletName string `mapstructure:"wallet-name"`
	// Destination defines the taproot transaction destination address
	Destination string `mapstructure:"destination"`
	// EnableIndexer defines whether to enable the indexer
	EnableIndexer bool `mapstructure:"enable-indexer"`
	// EnableCommitter defines whether to enable the committer
	EnableCommitter bool `mapstructure:"enable-committer"`
	// IndexerListenAddress defines the address to listen on
	IndexerListenAddress string `mapstructure:"indexer-listen-address"`
	// Bridge defines the bridge config
	Bridge BridgeConfig `mapstructure:"bridge"`
	// Dsn defines the state db dsn
	StateConfig StateConfig `mapstructure:"state"`
}

type BridgeConfig struct {
	EthRPCURL       string `mapstructure:"eth-rpc-url"`
	EthPrivKey      string `mapstructure:"eth-priv-key"`
	ContractAddress string `mapstructure:"contract-address"`
	ABI             string `mapstructure:"abi"`
	GasLimit        uint64 `mapstructure:"gas-limit"`
}

type StateConfig struct {
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	User   string `mapstructure:"user"`
	Pass   string `mapstructure:"pass"`
	DbName string `mapstructure:"db-name"`
}

const (
	BitcoinRPCConfigFileName = "bitcoin.toml"
)

func LoadBitcoinConfig(homePath string) (*BitconConfig, error) {
	config := BitconConfig{}
	configFile := path.Join(homePath, BitcoinRPCConfigFileName)
	_, err := os.Stat(configFile)
	if err != nil {
		// if file not exist use default config
		// TODO: add gen config command after, The default configuration may not be required
		if os.IsNotExist(err) {
			config.NetworkName = "signet"
			config.RPCHost = "localhost"
			config.RPCPort = "8332"
			config.RPCUser = "user"
			config.RPCPass = "password"
			config.WalletName = "walletname"
			config.Destination = "tb1qgm39cu009lyvq93afx47pp4h9wxq5x92lxxgnz"
			config.IndexerListenAddress = "tb1qsja4hvx66jr9grgmt8452letmz37gmludcrnup"
			config.EnableIndexer = false
			config.EnableCommitter = false
			return &config, nil
		}
		return nil, err
	}

	v := viper.New()
	v.SetConfigFile(configFile)
	v.AutomaticEnv()
	v.SetEnvPrefix("BITCOIN")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	err = v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// ChainParams get chain params by network name
func ChainParams(network string) *chaincfg.Params {
	switch network {
	case "mainnet":
		return &chaincfg.MainNetParams
	case "testnet":
		return &chaincfg.TestNet3Params
	case "signet":
		return &chaincfg.SigNetParams
	case "simnet":
		return &chaincfg.SimNetParams
	case "regtest":
		return &chaincfg.RegressionNetParams
	default:
		return &chaincfg.TestNet3Params
	}
}
