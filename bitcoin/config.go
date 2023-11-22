package bitcoin

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

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
	// IndexerListenAddress defines the address to listen on
	IndexerListenAddress string `mapstructure:"indexer-listen-address"`
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
			return &config, nil
		}
		return nil, err
	}

	v := viper.New()
	v.SetConfigFile(configFile)

	// TODO: set env prifix

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	err = v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
