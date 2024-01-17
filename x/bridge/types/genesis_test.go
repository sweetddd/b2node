package types_test

import (
	"testing"

	"github.com/evmos/ethermint/x/bridge/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				SignerGroupList: []types.SignerGroup{
					{
						Name: "0",
					},
					{
						Name: "1",
					},
				},
				CallerGroupList: []types.CallerGroup{
					{
						Name: "0",
					},
					{
						Name: "1",
					},
				},
				DepositList: []types.Deposit{
					{
						TxHash: "0",
					},
					{
						TxHash: "1",
					},
				},
				WithdrawList: []types.Withdraw{
					{
						TxHash: "0",
					},
					{
						TxHash: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated signerGroup",
			genState: &types.GenesisState{
				SignerGroupList: []types.SignerGroup{
					{
						Name: "0",
					},
					{
						Name: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated callerGroup",
			genState: &types.GenesisState{
				CallerGroupList: []types.CallerGroup{
					{
						Name: "0",
					},
					{
						Name: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated deposit",
			genState: &types.GenesisState{
				DepositList: []types.Deposit{
					{
						TxHash: "0",
					},
					{
						TxHash: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated withdraw",
			genState: &types.GenesisState{
				WithdrawList: []types.Withdraw{
					{
						TxHash: "0",
					},
					{
						TxHash: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
