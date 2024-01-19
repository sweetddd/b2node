package types_test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/evmos/ethermint/testutil"
	"github.com/evmos/ethermint/x/committer/types"
)

func TestMsgTimeoutPoposal(t *testing.T){
	addr := testutil.AccAddress()
	txs := []struct {
		name 							string
		msg 								types.MsgTimeoutProposal
		isError 					bool
		errMsg            string
	}{
		{
			name: "success",
			msg: types.MsgTimeoutProposal{
				From: addr,
				Id: 1,
			},
			isError: false,
		},
		{
			name: "failed with missing from address",
			msg: types.MsgTimeoutProposal{
				From: "",
				Id: 1,
			},
			isError: true,
			errMsg: "missing from address",
		},
		{
			name: "failed with invalid from address",
			msg: types.MsgTimeoutProposal{
				From: "invalid_address",
				Id: 1,
			},
			isError: true,
			errMsg: "invalid from address",
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