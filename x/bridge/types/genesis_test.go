package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"zyx/x/bridge/types"
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

				CallerList: []types.Caller{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				CallerCount: 2,
				DepositList: []types.Deposit{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				WithdrawList: []types.Withdraw{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				SignerList: []types.Signer{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				SignerCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated caller",
			genState: &types.GenesisState{
				CallerList: []types.Caller{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid caller count",
			genState: &types.GenesisState{
				CallerList: []types.Caller{
					{
						Id: 1,
					},
				},
				CallerCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated deposit",
			genState: &types.GenesisState{
				DepositList: []types.Deposit{
					{
						Index: "0",
					},
					{
						Index: "0",
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
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated signer",
			genState: &types.GenesisState{
				SignerList: []types.Signer{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid signer count",
			genState: &types.GenesisState{
				SignerList: []types.Signer{
					{
						Id: 1,
					},
				},
				SignerCount: 0,
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
