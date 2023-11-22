package bitcoin

import (
	"time"

	"github.com/tendermint/tendermint/libs/service"

	ethermint "github.com/evmos/ethermint/types"
)

const (
	ServiceName = "BitcoinIndexerService"

	NewBlockWaitTimeout = 1 * time.Second
)

// BitcoinIndexerService indexes transactions for json-rpc service.
type BitcoinIndexerService struct {
	service.BaseService

	txIdxr ethermint.BITCOINTxIndexer
}

// NewBitcoinIndexerService returns a new service instance.
func NewBitcoinIndexerService(
	txIdxr ethermint.BITCOINTxIndexer,
) *BitcoinIndexerService {
	is := &BitcoinIndexerService{txIdxr: txIdxr}
	is.BaseService = *service.NewBaseService(nil, ServiceName, is)
	return is
}

// OnStart
func (bis *BitcoinIndexerService) OnStart() error {
	latestBlock, err := bis.txIdxr.LatestBlock()
	if err != nil {
		bis.Logger.Error("bitcoin indexer latestBlock err", err.Error())
		return err
	}
	// TODO: load from kv store
	var currentBlock int64 = latestBlock

	ticker := time.NewTicker(NewBlockWaitTimeout)

	for {

		bis.Logger.Info("bitcoin indexer", "latestBlock", latestBlock, "currentIndexerBlock", currentBlock)

		if latestBlock <= currentBlock {
			<-ticker.C

			ticker.Reset(NewBlockWaitTimeout)

			// update latest block
			latestBlock, err = bis.txIdxr.LatestBlock()
			if err != nil {
				bis.Logger.Error("bitcoin indexer latestBlock err", err.Error())
			}

			continue
		}

		for i := currentBlock + 1; i <= latestBlock; i++ {
			txResult, err := bis.txIdxr.ParseBlock(i)
			if err != nil {
				bis.Logger.Error("bitcoin indexer parseblock err", err.Error())
				continue
			}

			//TODO: sleep  prevent frequent request to bitcoin core
			if i%10 == 0 {
				time.Sleep(500 * time.Millisecond)
			}
			if len(txResult) > 0 {
				bis.Logger.Info("bitcoin indexer parseblock success, send data", txResult)
				//TODO: send data
			}
			bis.Logger.Info("bitcoin indexer parsed", "txResult", txResult, "currentBlock", currentBlock, "latestBlock", latestBlock)
			currentBlock = i

		}
	}
}
