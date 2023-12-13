package bitcoin

import (
	"encoding/binary"
	dbm "github.com/tendermint/tm-db"
	"time"

	"github.com/tendermint/tendermint/libs/service"
)

const (
	BitcoinServiceName = "BitcoinCommitterService"

	WaitTimeout = 1 * time.Minute
)

// CommitterService is a service that commits bitcoin transactions.
type CommitterService struct {
	service.BaseService
	committer *Committer
	db        dbm.DB
}

// NewIndexerService returns a new service instance.
func NewCommitterService(
	committer *Committer,
	db dbm.DB,
) *CommitterService {
	is := &CommitterService{committer: committer}
	is.BaseService = *service.NewBaseService(nil, BitcoinServiceName, is)
	is.db = db
	return is
}

// OnStart
func (bis *CommitterService) OnStart() error {
	ticker := time.NewTicker(WaitTimeout)
	for {
		bis.Logger.Info("committer start....")
		<-ticker.C
		ticker.Reset(WaitTimeout)

		var index = int64(0)
		blockNumMax, err := bis.db.Get([]byte("blockNumMax"))
		if err != nil {
			bis.Logger.Error("Failed to get blockNumMax", "err", err)
			continue
		}
		if blockNumMax != nil {
			index = int64(binary.BigEndian.Uint64(blockNumMax))
		}

		roots, err := GetStateRoot(bis.committer.stateConfig, index)
		if roots == nil {
			continue
		}
		dataList := make([]InscriptionData, 0)
		for _, root := range roots {
			dataList = append(dataList, InscriptionData{
				Body:        []byte(root.StateRoot),
				Destination: bis.committer.destination,
			})
			if root.BlockNum > index {
				index = root.BlockNum
			}
		}

		req, err := NewRequest(bis.committer.client, dataList) // update latest block
		if err != nil {
			bis.Logger.Error("committer init req error", "err", err)
			continue
		}
		tool, err := NewInscriptionTool(bis.committer.chainParams, bis.committer.client, req)
		if err != nil {
			bis.Logger.Error("Failed to create inscription tool", "err", err)
			continue
		}
		err = tool.BackupRecoveryKeyToRPCNode()
		if err != nil {
			bis.Logger.Error("Failed to backup recovery key", "err", err)
			continue
		}
		commitTxHash, revealTxHashList, inscriptions, fees, err := tool.Inscribe()
		if err != nil {
			bis.Logger.Error("Failed to inscribe", "err", err)
			continue
		}
		bis.Logger.Info("commitTxHash," + commitTxHash.String())
		for i := range revealTxHashList {
			bis.Logger.Info("revealTxHash," + revealTxHashList[i].String())
		}
		for i := range inscriptions {
			bis.Logger.Info("inscriptions," + inscriptions[i])
		}
		bis.Logger.Info("fees:", "fee", fees)

		var buf = make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(index))
		err = bis.db.Set([]byte("blockNumMax"), buf)
		if err != nil {
			bis.Logger.Error("Failed to set blockNumMax", "err", err)
			continue
		}
	}
}
