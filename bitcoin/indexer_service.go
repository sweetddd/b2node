package bitcoin

import (
	"time"

	"github.com/tendermint/tendermint/libs/service"

	"github.com/evmos/ethermint/types"
)

const (
	ServiceName = "BitcoinIndexerService"

	NewBlockWaitTimeout = 1 * time.Second
)

// IndexerService indexes transactions for json-rpc service.
type IndexerService struct {
	service.BaseService

	txIdxr types.BITCOINTxIndexer
	bridge types.BITCOINBridge
}

// NewIndexerService returns a new service instance.
func NewIndexerService(
	txIdxr types.BITCOINTxIndexer,
	bridge types.BITCOINBridge,
) *IndexerService {
	is := &IndexerService{txIdxr: txIdxr, bridge: bridge}
	is.BaseService = *service.NewBaseService(nil, ServiceName, is)
	return is
}

// OnStart
func (bis *IndexerService) OnStart() error {
	latestBlock, err := bis.txIdxr.LatestBlock()
	if err != nil {
		bis.Logger.Error("bitcoin indexer latestBlock", "error", err.Error())
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
				bis.Logger.Error("bitcoin indexer latestBlock", "error", err.Error())
			}
			continue
		}

		for i := currentBlock + 1; i <= latestBlock; i++ {
			txResults, err := bis.txIdxr.ParseBlock(i)
			if err != nil {
				bis.Logger.Error("bitcoin indexer parseblock", "error", err.Error())
				continue
			}

			// TODO: sleep  prevent frequent request to bitcoin core
			if i%10 == 0 {
				time.Sleep(500 * time.Millisecond)
			}
			if len(txResults) > 0 {
				// TODO: temp test, Retries need to be considered
				for _, v := range txResults {
					if err := bis.bridge.Deposit(v.From[0], v.Value); err != nil {
						bis.Logger.Error("bitcoin indexer Deposit", "error", err.Error())
					}
				}
			}
			bis.Logger.Info("bitcoin indexer parsed", "txResult", txResults, "currentBlock", currentBlock, "latestBlock", latestBlock)
			currentBlock = i
		}
	}
}
