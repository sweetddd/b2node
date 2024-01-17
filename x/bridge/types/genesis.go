package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		SignerGroupList: []SignerGroup{},
		CallerGroupList: []CallerGroup{},
		DepositList:     []Deposit{},
		WithdrawList:    []Withdraw{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in signerGroup
	signerGroupIndexMap := make(map[string]struct{})

	for _, elem := range gs.SignerGroupList {
		index := string(SignerGroupKey(elem.Name))
		if _, ok := signerGroupIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for signerGroup")
		}
		signerGroupIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in callerGroup
	callerGroupIndexMap := make(map[string]struct{})

	for _, elem := range gs.CallerGroupList {
		index := string(CallerGroupKey(elem.Name))
		if _, ok := callerGroupIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for callerGroup")
		}
		callerGroupIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in deposit
	depositIndexMap := make(map[string]struct{})

	for _, elem := range gs.DepositList {
		index := string(DepositKey(elem.TxHash))
		if _, ok := depositIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for deposit")
		}
		depositIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in withdraw
	withdrawIndexMap := make(map[string]struct{})

	for _, elem := range gs.WithdrawList {
		index := string(WithdrawKey(elem.TxHash))
		if _, ok := withdrawIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for withdraw")
		}
		withdrawIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
