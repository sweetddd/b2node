package types_test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/evmos/ethermint/testutil"
	"github.com/evmos/ethermint/x/committer/types"
)

func TestMsgBitcoinTx(t *testing.T) {
	addr := testutil.AccAddress()

	txs := []struct {
		name 							string
		msg 								types.MsgBitcoinTx
		isError 					bool
		errMsg            string
	}{
		{
			name: "success",
			msg: types.MsgBitcoinTx{
				Id: 1,
				From: addr,
				BitcoinTxHash: "bitcoin_tx_hash",
			},
			isError: false,
		},
		{
			name: "failed with missing from address",
			msg: types.MsgBitcoinTx{
				Id: 1,
				From: "",
				BitcoinTxHash: "bitcoin_tx_hash",
			},
			isError: true,
			errMsg: "missing from address",
		},
		{
			name: "failed with invalid from address",
			msg: types.MsgBitcoinTx{
				Id: 1,
				From: "invalid_address",
				BitcoinTxHash: "bitcoin_tx_hash",
			},
			isError: true,
			errMsg: "invalid from address",
		},
		{
			name: "failed with missing bitcoin tx hash",
			msg: types.MsgBitcoinTx{
				Id: 1,
				From: addr,
				BitcoinTxHash: "",
			},
			isError: true,
			errMsg: "missing bitcoin tx hash",
		},
	}

	for _, tx := range txs {
		t.Run(tx.name, func(t *testing.T) {
			err := tx.msg.ValidateBasic()
			if tx.isError {
				require.Error(t, err)
				require.Contains(t, err.Error(), tx.errMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}