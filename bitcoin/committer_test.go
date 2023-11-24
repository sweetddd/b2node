package bitcoin_test

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
	"github.com/evmos/ethermint/bitcoin"
	"log"
	"testing"
)

func TestNewInscriptionTool(t *testing.T) {
	config, err := bitcoin.LoadBitcoinConfig("./testdata")
	connCfg := &rpcclient.ConnConfig{
		Host:         config.RPCHost + ":" + config.RPCPort + "/wallet/" + config.WalletName,
		User:         config.RPCUser,
		Pass:         config.RPCPass,
		HTTPPostMode: true,
		DisableTLS:   true,
	}
	client, _ := rpcclient.New(connCfg, nil)
	commitTxOutPointList := make([]*wire.OutPoint, 0)
	unspentList, _ := client.ListUnspent()
	for i := range unspentList {
		inTxid, _ := chainhash.NewHashFromStr(unspentList[i].TxID)
		commitTxOutPointList = append(commitTxOutPointList, wire.NewOutPoint(inTxid, unspentList[i].Vout))
	}
	req := bitcoin.InscriptionRequest{
		CommitTxOutPointList: commitTxOutPointList,
		CommitFeeRate:        25,
		FeeRate:              26,
		DataList: []bitcoin.InscriptionData{
			{
				Destination: config.Destination,
				Body:        []byte("hello world"),
			},
		},
		RevealOutValue: 500,
	}
	tool, err := bitcoin.NewInscriptionTool(&chaincfg.TestNet3Params, client, &req)
	if err != nil {
		log.Fatalf("Failed to create inscription tool: %v", err)
	}
	err = tool.BackupRecoveryKeyToRPCNode()
	if err != nil {
		log.Fatalf("Failed to backup recovery key: %v", err)
	}
	commitTxHash, revealTxHashList, inscriptions, fees, err := tool.Inscribe()
	if err != nil {
		log.Fatalf("send tx errr, %v", err)
	}
	log.Println("commitTxHash, " + commitTxHash.String())
	for i := range revealTxHashList {
		log.Println("revealTxHash, " + revealTxHashList[i].String())
	}
	for i := range inscriptions {
		log.Println("inscription, " + inscriptions[i])
	}
	log.Println("fees: ", fees)
}
