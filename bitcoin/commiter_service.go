package bitcoin

import (
	"math/rand"
	"time"

	"github.com/tendermint/tendermint/libs/service"
)

const (
	BitcoinServiceName = "BitcoinCommiterService"

	WaitTimeout = 1 * time.Minute
)

// CommitterService is a service that commits bitcoin transactions.
type CommitterService struct {
	service.BaseService
	committer *Committer
}

// NewIndexerService returns a new service instance.
func NewCommitterService(
	committer *Committer,
) *CommitterService {
	is := &CommitterService{committer: committer}
	is.BaseService = *service.NewBaseService(nil, BitcoinServiceName, is)
	return is
}

// OnStart
func (bis *CommitterService) OnStart() error {
	ticker := time.NewTicker(WaitTimeout)
	for {
		bis.Logger.Info("committer start....")
		<-ticker.C
		ticker.Reset(WaitTimeout)

		dataList := make([]InscriptionData, 0)

		dataList = append(dataList, InscriptionData{
			Body:        []byte(RandomStr(16)),
			Destination: bis.committer.destination,
		})

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
		err = tool.BackupRecoveryKeyToRpcNode()
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
	}
}

// RandomStr generate random string
func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
