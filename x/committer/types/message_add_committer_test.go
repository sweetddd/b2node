package types_test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/evmos/ethermint/testutil"
	"github.com/evmos/ethermint/x/committer/types"
)

func TestMsgAddCommitter(t *testing.T) {
	addr := testutil.AccAddress()
	txs := []struct {
		name 							string
		msg 								types.MsgAddCommitter
		isError 					bool
		errMsg            string
	}{
		{
			name: "success",
			msg: types.MsgAddCommitter{
				From: addr,
				Committer: addr,
			},
			isError: false,
		},
		{
			name: "failed with missing from address",
			msg: types.MsgAddCommitter{
				From: "",
				Committer: addr,
			},
			isError: true,
			errMsg: "missing from address",
		},
		{
			name: "failed with invalid from address",
			msg: types.MsgAddCommitter{
				From: "invalid_address",
				Committer: addr,
			},
			isError: true,
			errMsg: "invalid from address",
		},
		{
			name: "failed with missing committer",
			msg: types.MsgAddCommitter{
				From: addr,
				Committer: "",
			},
			isError: true,
			errMsg: "missing committer",
		},
		{
			name: "failed with invalid committer",
			msg: types.MsgAddCommitter{
				From: addr,
				Committer: "invalid_address",
			},
			isError: true,
			errMsg: "invalid committer address",
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