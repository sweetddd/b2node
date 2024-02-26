package types_test

import (
	"testing"

	"github.com/evmos/ethermint/testutil"
	"github.com/evmos/ethermint/x/committer/types"
	"github.com/stretchr/testify/require"
)

func TestMsgSubmitProof(t *testing.T) {
	addr := testutil.AccAddress()
	txs := []struct {
		name    string
		msg     types.MsgSubmitProof
		isError bool
		errMsg  string
	}{
		{
			name: "success",
			msg: types.MsgSubmitProof{
				Id:            1,
				From:          addr,
				ProofHash:     "proof",
				StateRootHash: "state_root_hash",
				StartIndex:    1,
				EndIndex:      2,
			},
			isError: false,
		},
		{
			name: "failed with missing from address",
			msg: types.MsgSubmitProof{
				Id:            1,
				From:          "",
				ProofHash:     "proof",
				StateRootHash: "state_root_hash",
				StartIndex:    1,
				EndIndex:      2,
			},
			isError: true,
			errMsg:  "missing from address",
		},
		{
			name: "failed with invalid from address",
			msg: types.MsgSubmitProof{
				Id:            1,
				From:          "invalid_address",
				ProofHash:     "proof",
				StateRootHash: "state_root_hash",
				StartIndex:    1,
				EndIndex:      2,
			},
			isError: true,
			errMsg:  "invalid from address",
		},
		{
			name: "failed with missing proof hash",
			msg: types.MsgSubmitProof{
				Id:            1,
				From:          addr,
				ProofHash:     "",
				StateRootHash: "state_root_hash",
				StartIndex:    1,
				EndIndex:      2,
			},
			isError: true,
			errMsg:  "missing proof",
		},
		{
			name: "failed with missing state root hash",
			msg: types.MsgSubmitProof{
				Id:            1,
				From:          addr,
				ProofHash:     "proof_hash",
				StateRootHash: "",
				StartIndex:    1,
				EndIndex:      2,
			},
			isError: true,
		},
		{
			name: "failed with start index greater than end index",
			msg: types.MsgSubmitProof{
				Id:            1,
				From:          addr,
				ProofHash:     "proof_hash",
				StateRootHash: "state_root_hash",
				StartIndex:    2,
				EndIndex:      1,
			},
			isError: true,
			errMsg:  "start index must be less than end index",
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
