package types_test

import (
	"testing"

	"github.com/evmos/ethermint/testutil"
	"github.com/evmos/ethermint/x/committer/types"
	"github.com/stretchr/testify/require"
)

func TestMsgRemoveCommitter(t *testing.T) {
	addr := testutil.AccAddress()
	txs := []struct {
		name    string
		msg     types.MsgRemoveCommitter
		isError bool
		errMsg  string
	}{
		{
			name: "success",
			msg: types.MsgRemoveCommitter{
				From:      addr,
				Committer: addr,
			},
			isError: false,
		},
		{
			name: "failed with missing from address",
			msg: types.MsgRemoveCommitter{
				From:      "",
				Committer: addr,
			},
			isError: true,
			errMsg:  "missing from address",
		},
		{
			name: "failed with invalid from address",
			msg: types.MsgRemoveCommitter{
				From:      "invalid_address",
				Committer: addr,
			},
			isError: true,
			errMsg:  "invalid from address",
		},
		{
			name: "failed with missing committer",
			msg: types.MsgRemoveCommitter{
				From:      addr,
				Committer: "",
			},
			isError: true,
			errMsg:  "missing committer",
		},
		{
			name: "failed with invalid committer",
			msg: types.MsgRemoveCommitter{
				From:      addr,
				Committer: "invalid_address",
			},
			isError: true,
			errMsg:  "invalid committer address",
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
