package bitcoin_test

import (
	"testing"

	"github.com/btcsuite/btcd/rpcclient"

	"github.com/stretchr/testify/require"
)

func TestNewCommitterRpcClient(t *testing.T) {
	connCfg := &rpcclient.ConnConfig{
		Host:         "43.157.135.191:8332/wallet/test",
		User:         "b2node",
		Pass:         "b2node@pass",
		HTTPPostMode: true,
		DisableTLS:   true,
	}
	client, err := rpcclient.New(connCfg, nil)
	require.NoError(t, err)
	require.NotEmpty(t, client)
}

//func TestNewInscriptionTool(t *testing.T) {
//	connCfg := &rpcclient.ConnConfig{
//		Host:         "43.157.135.191:8332/wallet/test",
//		User:         "b2node",
//		Pass:         "b2node@pass",
//		HTTPPostMode: true,
//		DisableTLS:   true,
//	}
//	client, _ := rpcclient.New(connCfg, nil)
//	commitTxOutPointList := make([]*wire.OutPoint, 0)
//	unspentList, _ := client.ListUnspent()
//	for i := range unspentList {
//		inTxid, _ := chainhash.NewHashFromStr(unspentList[i].TxID)
//		commitTxOutPointList = append(commitTxOutPointList, wire.NewOutPoint(inTxid, unspentList[i].Vout))
//	}
//	req := bitcoin.InscriptionRequest{
//		CommitTxOutPointList: commitTxOutPointList,
//		CommitFeeRate:        25,
//		FeeRate:              26,
//		DataList: []bitcoin.InscriptionData{
//			{
//				Destination: "tb1q9j03nm97urq4vwkt3mhfh2hgfgwvq329yekdc2",
//				Body:        []byte(bitcoin.RandomStr(16)),
//			},
//		},
//		RevealOutValue: 500,
//	}
//	tool, err := bitcoin.NewInscriptionTool(&chaincfg.TestNet3Params, client, &req)
//	if err != nil {
//		log.Fatalf("Failed to create inscription tool: %v", err)
//	}
//	err = tool.BackupRecoveryKeyToRpcNode()
//	if err != nil {
//		log.Fatalf("Failed to backup recovery key: %v", err)
//	}
//	commitTxHash, revealTxHashList, inscriptions, fees, err := tool.Inscribe()
//	if err != nil {
//		log.Fatalf("send tx errr, %v", err)
//	}
//	log.Println("commitTxHash, " + commitTxHash.String())
//	for i := range revealTxHashList {
//		log.Println("revealTxHash, " + revealTxHashList[i].String())
//	}
//	for i := range inscriptions {
//		log.Println("inscription, " + inscriptions[i])
//	}
//	log.Println("fees: ", fees)
//}
