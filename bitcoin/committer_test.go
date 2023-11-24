package bitcoin_test

import (
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewCommitterRpcClient(t *testing.T) {
	viper := viper.New()
	viper.SetEnvPrefix("bitcoin")
	viper.AutomaticEnv()
	//viper.SetEnvKeyReplacer(strings.NewReplacer("_", "-"))
	host := viper.GetString("rpc_host")
	port := viper.GetString("rpc_port")
	walletName := viper.GetString("wallet_name")
	rpcUser := viper.GetString("rpc_user")
	rpcPass := viper.GetString("rpc_pass")

	connCfg := &rpcclient.ConnConfig{
		Host:         host + ":" + port + "/wallet/" + walletName,
		User:         rpcUser,
		Pass:         rpcPass,
		HTTPPostMode: true,
		DisableTLS:   true,
	}
	client, err := rpcclient.New(connCfg, nil)
	require.NoError(t, err)
	require.NotEmpty(t, client)
}
