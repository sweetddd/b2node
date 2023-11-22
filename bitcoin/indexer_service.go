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

// IndexerService indexes transactions for json-rpc service.
type IndexerService struct {
	service.BaseService

	txIdxr ethermint.BITCOINTxIndexer
}

// NewIndexerService returns a new service instance.
func NewIndexerService(
	txIdxr ethermint.BITCOINTxIndexer,
) *IndexerService {
	is := &IndexerService{txIdxr: txIdxr}
	is.BaseService = *service.NewBaseService(nil, ServiceName, is)
	return is
}

// OnStart
func (bis *IndexerService) OnStart() error {
	latestBlock, err := bis.txIdxr.LatestBlock()
	if err != nil {
		bis.Logger.Error("bitcoin indexer latestBlock err", err.Error())
		return err
	}
	// TODO: load from kv store
	currentBlock := latestBlock

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

			// TODO: sleep  prevent frequent request to bitcoin core
			if i%10 == 0 {
				time.Sleep(500 * time.Millisecond)
			}
			if len(txResult) > 0 {
				bis.Logger.Info("bitcoin indexer parseblock success, send data", txResult)
				// TODO: send data
			}
			bis.Logger.Info("bitcoin indexer parsed", "txResult", txResult, "currentBlock", currentBlock, "latestBlock", latestBlock)
			currentBlock = i
		}
	}
}
