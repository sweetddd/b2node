package bitcoin

var BITCOIN_RPC BitcoinClient

type BitcoinClient struct {
}

func New() BitcoinClient {
	BITCOIN_RPC = BitcoinClient{}
	return BITCOIN_RPC
}
